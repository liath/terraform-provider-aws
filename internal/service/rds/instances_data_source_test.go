// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rds_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/rds"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccRDSInstancesDataSource_filter(t *testing.T) {
	ctx := acctest.Context(t)
	var dbInstance rds.DBInstance
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_db_instances.test"
	resourceName := "aws_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccInstancesDataSourceConfig_filter(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInstanceExists(ctx, resourceName, &dbInstance),
					resource.TestCheckResourceAttr(dataSourceName, "instance_arns.#", "1"),
					resource.TestCheckResourceAttrPair(dataSourceName, "instance_arns.0", resourceName, "arn"),
					resource.TestCheckResourceAttr(dataSourceName, "instance_identifiers.#", "1"),
					resource.TestCheckResourceAttrPair(dataSourceName, "instance_identifiers.0", resourceName, "identifier"),
				),
			},
		},
	})
}

func TestAccRDSInstancesDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var dbInstance rds.DBInstance
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	// Resources.
	resourceInstanceBlueStaging := "aws_db_instance.blue_staging"
	resourceInstanceGreenStaging := "aws_db_instance.green_staging"
	// Data sources.
	datasourceBlueTeamInstance := "data.aws_db_instances.get_instance_of_blue_team"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckInstanceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccInstancesDataSourceConfig_tag(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInstanceExists(ctx, resourceInstanceBlueStaging, &dbInstance),
					testAccCheckInstanceExists(ctx, resourceInstanceGreenStaging, &dbInstance),
					resource.TestCheckResourceAttrPair(datasourceBlueTeamInstance, "instance_arns.0", resourceInstanceBlueStaging, "arn"),
				),
			},
		},
	})
}

func testAccInstancesDataSourceConfig_filter(rName string) string {
	return fmt.Sprintf(`
data "aws_rds_engine_version" "default" {
  engine = "postgres"
}

resource "aws_db_instance" "test" {
  identifier           = %[1]q
  allocated_storage    = 10
  engine               = data.aws_rds_engine_version.default.engine
  engine_version       = data.aws_rds_engine_version.default.version
  instance_class       = "db.t4g.micro"
  db_name              = "test"
  password             = "avoid-plaintext-passwords"
  username             = "tfacctest"
  parameter_group_name = "default.${data.aws_rds_engine_version.default.parameter_group_family}"
  skip_final_snapshot  = true

  apply_immediately = true
}

resource "aws_db_instance" "wrong" {
  identifier           = "wrong-%[1]s"
  allocated_storage    = 10
  engine               = data.aws_rds_engine_version.default.engine
  engine_version       = data.aws_rds_engine_version.default.version
  instance_class       = "db.t4g.micro"
  db_name              = "test"
  password             = "avoid-plaintext-passwords"
  username             = "tfacctest"
  parameter_group_name = "default.${data.aws_rds_engine_version.default.parameter_group_family}"
  skip_final_snapshot  = true

  apply_immediately = true
}


data "aws_db_instances" "test" {
  filter {
    name   = "db-instance-id"
    values = [aws_db_instance.test.identifier]
  }
}
`, rName)
}

func testAccInstancesDataSourceConfig_tag(rName string) string {
	return fmt.Sprintf(`
data "aws_rds_engine_version" "default" {
  engine = "postgres"
}

resource "aws_db_instance" "blue_staging" {
  identifier           = "blue-staging-%[1]s"
  allocated_storage    = 10
  engine               = data.aws_rds_engine_version.default.engine
  engine_version       = data.aws_rds_engine_version.default.version
  instance_class       = "db.t4g.micro"
  db_name              = "test"
  password             = "avoid-plaintext-passwords"
  username             = "tfacctest"
  parameter_group_name = "default.${data.aws_rds_engine_version.default.parameter_group_family}"
  skip_final_snapshot  = true

  tags = {
    Name        = "blue-staging-%[1]s",
    CostCenter  = "ResearchDepartment",
    Team        = "Blue",
    Environment = "staging"
  }

  apply_immediately = true
}

resource "aws_db_instance" "green_staging" {
  identifier           = "green-staging-%[1]s"
  allocated_storage    = 10
  engine               = data.aws_rds_engine_version.default.engine
  engine_version       = data.aws_rds_engine_version.default.version
  instance_class       = "db.t4g.micro"
  db_name              = "test"
  password             = "avoid-plaintext-passwords"
  username             = "tfacctest"
  parameter_group_name = "default.${data.aws_rds_engine_version.default.parameter_group_family}"
  skip_final_snapshot  = true

  tags = {
    Name        = "green-staging-%[1]s",
    CostCenter  = "ResearchDepartment",
    Team        = "Green",
    Environment = "staging"
  }

  apply_immediately = true
}


data "aws_db_instances" "get_instance_of_blue_team" {
  tag {
    key   = "Team"
    value = "Blue"
  }

  depends_on = [aws_db_instance.green_staging, aws_db_instance.blue_staging]
}
`, rName)
}
