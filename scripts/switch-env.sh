#!/bin/bash

# Environment switcher script for Bravo Services
# Usage: ./scripts/switch-env.sh [local|dev|prod]

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_color() {
    printf "${1}${2}${NC}\n"
}

# Function to show usage
show_usage() {
    print_color $BLUE "Usage: $0 [local|dev|prod]"
    print_color $BLUE "  local - Switch to local environment"
    print_color $BLUE "  dev   - Switch to development environment"
    print_color $BLUE "  prod  - Switch to production environment"
    print_color $BLUE ""
    print_color $BLUE "Current environment: $(get_current_env)"
}

# Function to get current environment
get_current_env() {
    if [ -f "$PROJECT_ROOT/.env" ]; then
        grep "^APP_ENV=" "$PROJECT_ROOT/.env" 2>/dev/null | cut -d'=' -f2 || echo "unknown"
    else
        echo "none"
    fi
}

# Function to create .env from template
create_env_from_template() {
    local env=$1
    local template_file="$PROJECT_ROOT/.env.$env"
    local target_file="$PROJECT_ROOT/.env"
    
    if [ -f "$template_file" ]; then
        cp "$template_file" "$target_file"
        print_color $GREEN "✓ Switched to $env environment"
        print_color $YELLOW "⚠ Remember to update your database credentials in .env if needed"
    else
        print_color $RED "✗ Environment file .env.$env not found"
        print_color $YELLOW "Creating from .env.example..."
        if [ -f "$PROJECT_ROOT/.env.example" ]; then
            cp "$PROJECT_ROOT/.env.example" "$template_file"
            # Set the APP_ENV in the new file
            sed -i.bak "s/^APP_ENV=.*/APP_ENV=$env/" "$template_file" 2>/dev/null || true
            cp "$template_file" "$target_file"
            print_color $GREEN "✓ Created and switched to $env environment"
            print_color $YELLOW "⚠ Please edit .env.$env and .env with your specific configuration"
        else
            print_color $RED "✗ .env.example not found. Cannot create environment file."
            exit 1
        fi
    fi
}

# Function to show current configuration
show_config() {
    local current_env=$(get_current_env)
    print_color $BLUE "Current Environment Configuration:"
    print_color $BLUE "================================="
    if [ -f "$PROJECT_ROOT/.env" ]; then
        grep -E "^(APP_ENV|POSTGRES_|PORT|GIN_MODE)=" "$PROJECT_ROOT/.env" | while read line; do
            print_color $GREEN "$line"
        done
    else
        print_color $RED "No .env file found"
    fi
}

# Main script logic
case "${1:-}" in
    "local"|"dev"|"prod")
        ENV=$1
        print_color $BLUE "Switching to $ENV environment..."
        create_env_from_template $ENV
        show_config
        ;;
    "status"|"current")
        show_config
        ;;
    "help"|"-h"|"--help")
        show_usage
        ;;
    "")
        print_color $YELLOW "No environment specified."
        show_usage
        ;;
    *)
        print_color $RED "✗ Invalid environment: $1"
        show_usage
        exit 1
        ;;
esac
