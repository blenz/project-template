locals {
  name = "test-auth"
}

provider "aws" {
  region = "us-west-2"
}

resource "aws_cognito_user_pool" "this" {
  name = "${local.name}-user-pool"

  password_policy {
    minimum_length = 8
  }

  admin_create_user_config {
    allow_admin_create_user_only = false
  }


  alias_attributes = ["preferred_username"]
}

resource "aws_cognito_user_pool_client" "this" {
  name                         = "${local.name}-client"
  user_pool_id                 = aws_cognito_user_pool.this.id
  generate_secret              = true
  supported_identity_providers = ["COGNITO"]

  allowed_oauth_flows_user_pool_client = true
  allowed_oauth_flows                  = ["code"]
  allowed_oauth_scopes                 = ["openid", "email", "profile"]

  callback_urls = ["http://localhost:3000/api/auth/callback"]
  logout_urls   = ["http://localhost:3000/api/auth/logout"]
}

resource "aws_cognito_user_pool_domain" "this" {
  domain       = "${local.name}-user-pool"
  user_pool_id = aws_cognito_user_pool.this.id
}
