# Provider configuration
provider "google" {
  project = "praxis-surface-441613-q1"
  region  = "us-central1"
}

# MongoDB instance
resource "google_compute_instance" "mongodb" {
  name         = "mongodb"
  machine_type = "e2-medium"
  zone         = "us-central1-b"

  # Boot disk configuration with Debian OS
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  # Network configuration
  network_interface {
    network = "default"
    access_config {}
  }

  # Startup script to disable sudo password prompt
  metadata_startup_script = <<-EOT
    #!/bin/bash
    echo "$(whoami) ALL=(ALL) NOPASSWD:ALL" | sudo tee /etc/sudoers.d/$(whoami)
  EOT

  tags = ["mongodb"]
}

# SSH Key configuration
resource "tls_private_key" "ssh_key" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "google_compute_project_metadata_item" "ssh_key" {
  key   = "ssh-keys"
  value = "hardik20507:${tls_private_key.ssh_key.public_key_openssh}"
}

# Remote-exec provisioner to install MongoDB server, backup data, and upload to Google Cloud Storage
resource "null_resource" "mongodb_provision" {
  depends_on = [google_compute_instance.mongodb]

  provisioner "remote-exec" {
    connection {
      type        = "ssh"
      user        = "hardik20507"
      private_key = tls_private_key.ssh_key.private_key_pem
      host        = google_compute_instance.mongodb.network_interface[0].access_config[0].nat_ip
    }

    inline = [
      # Retry logic for apt-get commands to avoid lock issues
      "for i in {1..5}; do sudo apt-get update && break || sleep 10; done",
      # Add MongoDB repository
      "wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add -",
      "echo 'deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/debian buster/mongodb-org/5.0 main' | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list",
      "for i in {1..5}; do sudo apt-get update && break || sleep 10; done",
      "for i in {1..5}; do sudo apt-get install -y mongodb-org && break || sleep 10; done",
      # Start and enable MongoDB service
      "sudo systemctl start mongod",
      "sudo systemctl enable mongod",
      # Extended delay for MongoDB to initialize
      "sleep 30",
      # Check if MongoDB is running
      "sudo systemctl is-active mongod || (echo 'MongoDB failed to start' && exit 1)",
      # Display MongoDB logs to confirm service status
      "sudo journalctl -u mongod --no-pager | tail -n 20",
      # Attempt to connect once to MongoDB
      "mongo --host localhost --quiet --eval \"db.runCommand('ping').ok\" || echo 'MongoDB is not yet responding'"
    ]
  }
}

# Attach block storage
resource "google_compute_disk" "mongodb_data" {
  name  = "mongodb-data"
  type  = "pd-standard"
  zone  = google_compute_instance.mongodb.zone
  size  = 50 # in GB
}

resource "google_compute_attached_disk" "mongodb_data_attachment" {
  disk     = google_compute_disk.mongodb_data.self_link
  instance = google_compute_instance.mongodb.self_link
}

# Backup solution
resource "google_storage_bucket" "mongodb_backup" {
  name     = "my-unique-mongodb-backup-bucket"
  location = "us-central1"
}

# Firewall rules
resource "google_compute_firewall" "allow_mongodb" {
  name    = "my-unique-mongodb-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["27017"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["mongodb"]
}

# Restore Resource with Retry Logic
resource "null_resource" "mongodb_restore" {
  depends_on = [null_resource.mongodb_provision]

  provisioner "remote-exec" {
    connection {
      type        = "ssh"
      user        = "hardik20507"
      private_key = tls_private_key.ssh_key.private_key_pem
      host        = google_compute_instance.mongodb.network_interface[0].access_config[0].nat_ip
    }

    inline = [
      # Retry logic for apt-get commands to avoid lock issues
      "for i in {1..5}; do sudo apt-get update && break || sleep 10; done",
      # Add MongoDB repository
      "wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add -",
      "echo 'deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/debian buster/mongodb-org/5.0 main' | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list",
      "for i in {1..5}; do sudo apt-get update && break || sleep 10; done",
      "for i in {1..5}; do sudo apt-get install -y mongodb-org && break || sleep 10; done",
      # Start and enable MongoDB service
      "sudo systemctl start mongod",
      "sudo systemctl enable mongod",
      # Extended delay for MongoDB to initialize
      "sleep 30",
      # Check if MongoDB is running
      "sudo systemctl is-active mongod || (echo 'MongoDB failed to start' && exit 1)",
      # Display MongoDB logs to confirm service status
      "sudo journalctl -u mongod --no-pager | tail -n 20",
      # Attempt to connect once to MongoDB
      "mongo --host localhost --quiet --eval \"db.runCommand('ping').ok\" || echo 'MongoDB is not yet responding'",
      # Restore MongoDB data from backup if available
      "if gsutil -q stat gs://my-unique-mongodb-backup-bucket/latest.tar.gz; then gsutil cp gs://my-unique-mongodb-backup-bucket/latest.tar.gz /tmp/mongodb_backup.tar.gz && mongorestore --host localhost --gzip --archive=/tmp/mongodb_backup.tar.gz; else echo 'No backup found to restore'; fi"
    ]
  }
}
