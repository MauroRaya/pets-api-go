terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.0.0"
    }
  }
}

provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "pets_rg" {
  name     = "pets-rg"
  location = "Canada Central"
}

resource "azurerm_container_registry" "pets_acr" {
  name                = "petsacr987"
  location            = azurerm_resource_group.pets_rg.location
  resource_group_name = azurerm_resource_group.pets_rg.name
  sku                 = "Basic"
  admin_enabled       = true
}

resource "azurerm_app_service_plan" "pets_plan" {
  name                = "pets-service-plan"
  location            = azurerm_resource_group.pets_rg.location
  resource_group_name = azurerm_resource_group.pets_rg.name
  kind                = "Linux"

  sku {
    tier = "Free"
    size = "F1"
  }
}

resource "azurerm_app_service" "pets_api" {
  name                = "pets-api"
  location            = azurerm_resource_group.pets_rg.location
  resource_group_name = azurerm_resource_group.pets_rg.name
  app_service_plan_id = azurerm_app_service_plan.pets_plan.id

  site_config {
    linux_fx_version = "DOCKER|pets_image:latest"
  }

  app_settings = {
    WEBSITES_ENABLE_APP_SERVICE_STORAGE = "false"
    DOCKER_ENABLE_CI                    = "true"
  }
}