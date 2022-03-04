terraform {
  backend "remote" {
    organization = "hashicorp-v2"

    workspaces {
      name = "terraform-provider-aws-repository"
    }
  }

  required_providers {
    github = {
      source  = "hashicorp/github"
      version = "4.20.1"
    }
  }

  required_version = ">= 0.13.5"
}

provider "github" {
  organization = "hashicorp"
}
