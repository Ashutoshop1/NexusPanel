#!/bin/bash

# NexusPanel Installation Script
# This script installs NexusPanel on Linux systems

set -e

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Print colored messages
print_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if running as root
if [ "$EUID" -ne 0 ]; then 
    print_error "Please run as root or with sudo"
    exit 1
fi

print_info "Starting NexusPanel installation..."

# Detect OS
if [ -f /etc/os-release ]; then
    . /etc/os-release
    OS=$ID
    VER=$VERSION_ID
else
    print_error "Unable to detect operating system"
    exit 1
fi

print_info "Detected OS: $OS $VER"

# Install Docker if not present
if ! command -v docker &> /dev/null; then
    print_info "Docker not found, installing..."
    curl -fsSL https://get.docker.com | sh
    systemctl enable docker
    systemctl start docker
    print_info "Docker installed successfully"
else
    print_info "Docker is already installed"
fi

# Install Docker Compose if not present
if ! command -v docker-compose &> /dev/null; then
    print_info "Docker Compose not found, installing..."
    curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose
    print_info "Docker Compose installed successfully"
else
    print_info "Docker Compose is already installed"
fi

# Create installation directory
INSTALL_DIR="/opt/nexuspanel"
print_info "Creating installation directory: $INSTALL_DIR"
mkdir -p $INSTALL_DIR
cd $INSTALL_DIR

# Download docker-compose.yml
print_info "Downloading docker-compose.yml..."
curl -fsSL https://raw.githubusercontent.com/2670044605/NexusPanel/main/deploy/docker/docker-compose.yml -o docker-compose.yml

# Download example config
print_info "Downloading configuration file..."
curl -fsSL https://raw.githubusercontent.com/2670044605/NexusPanel/main/configs/config.example.yaml -o config.yaml

# Generate random passwords
POSTGRES_PASSWORD=$(openssl rand -base64 32)
JWT_SECRET=$(openssl rand -base64 32)
SESSION_SECRET=$(openssl rand -base64 32)
SSH_ENCRYPTION_KEY=$(openssl rand -base64 32)

# Update config with generated secrets
print_info "Configuring with secure random secrets..."
sed -i "s/your-password/$POSTGRES_PASSWORD/g" config.yaml
sed -i "s/change-this-jwt-secret-in-production/$JWT_SECRET/g" config.yaml
sed -i "s/change-this-secret-key-in-production/$SESSION_SECRET/g" config.yaml
sed -i "s/change-this-encryption-key-32-chars-long!!/$SSH_ENCRYPTION_KEY/g" config.yaml

# Update docker-compose with generated password
sed -i "s/nexuspanel123/$POSTGRES_PASSWORD/g" docker-compose.yml

# Create systemd service
print_info "Creating systemd service..."
cat > /etc/systemd/system/nexuspanel.service <<EOF
[Unit]
Description=NexusPanel Server Management Platform
Requires=docker.service
After=docker.service

[Service]
Type=oneshot
RemainAfterExit=yes
WorkingDirectory=$INSTALL_DIR
ExecStart=/usr/local/bin/docker-compose up -d
ExecStop=/usr/local/bin/docker-compose down
TimeoutStartSec=0

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd
systemctl daemon-reload

# Start NexusPanel
print_info "Starting NexusPanel..."
systemctl enable nexuspanel
systemctl start nexuspanel

# Wait for services to be ready
print_info "Waiting for services to start..."
sleep 10

# Check if services are running
if systemctl is-active --quiet nexuspanel; then
    print_info "✓ NexusPanel is running"
else
    print_error "✗ NexusPanel failed to start"
    print_info "Check logs with: journalctl -u nexuspanel -f"
    exit 1
fi

# Get server IP
SERVER_IP=$(hostname -I | awk '{print $1}')

# Print success message
echo ""
echo "============================================"
print_info "NexusPanel installed successfully!"
echo "============================================"
echo ""
echo "Access NexusPanel at: http://$SERVER_IP:8080"
echo ""
echo "Default credentials:"
echo "  Username: admin"
echo "  Password: admin123"
echo ""
echo "⚠️  IMPORTANT: Change the default password after first login!"
echo ""
echo "Installation directory: $INSTALL_DIR"
echo "Configuration file: $INSTALL_DIR/config.yaml"
echo ""
echo "Useful commands:"
echo "  Start:   systemctl start nexuspanel"
echo "  Stop:    systemctl stop nexuspanel"
echo "  Restart: systemctl restart nexuspanel"
echo "  Status:  systemctl status nexuspanel"
echo "  Logs:    docker-compose -f $INSTALL_DIR/docker-compose.yml logs -f"
echo ""
echo "For more information, visit: https://github.com/2670044605/NexusPanel"
echo "============================================"
