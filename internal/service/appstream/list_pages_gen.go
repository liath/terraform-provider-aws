// Code generated by "internal/generate/listpages/main.go -ListOps=DescribeDirectoryConfigs,DescribeFleets,DescribeImageBuilders,DescribeStacks,DescribeUsers,ListAssociatedStacks"; DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appstream"
)

func describeDirectoryConfigsPages(conn *appstream.AppStream, input *appstream.DescribeDirectoryConfigsInput, fn func(*appstream.DescribeDirectoryConfigsOutput, bool) bool) error {
	return describeDirectoryConfigsPagesWithContext(context.Background(), conn, input, fn)
}

func describeDirectoryConfigsPagesWithContext(ctx context.Context, conn *appstream.AppStream, input *appstream.DescribeDirectoryConfigsInput, fn func(*appstream.DescribeDirectoryConfigsOutput, bool) bool) error {
	for {
		output, err := conn.DescribeDirectoryConfigsWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func describeFleetsPages(conn *appstream.AppStream, input *appstream.DescribeFleetsInput, fn func(*appstream.DescribeFleetsOutput, bool) bool) error {
	return describeFleetsPagesWithContext(context.Background(), conn, input, fn)
}

func describeFleetsPagesWithContext(ctx context.Context, conn *appstream.AppStream, input *appstream.DescribeFleetsInput, fn func(*appstream.DescribeFleetsOutput, bool) bool) error {
	for {
		output, err := conn.DescribeFleetsWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func describeImageBuildersPages(conn *appstream.AppStream, input *appstream.DescribeImageBuildersInput, fn func(*appstream.DescribeImageBuildersOutput, bool) bool) error {
	return describeImageBuildersPagesWithContext(context.Background(), conn, input, fn)
}

func describeImageBuildersPagesWithContext(ctx context.Context, conn *appstream.AppStream, input *appstream.DescribeImageBuildersInput, fn func(*appstream.DescribeImageBuildersOutput, bool) bool) error {
	for {
		output, err := conn.DescribeImageBuildersWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func describeStacksPages(conn *appstream.AppStream, input *appstream.DescribeStacksInput, fn func(*appstream.DescribeStacksOutput, bool) bool) error {
	return describeStacksPagesWithContext(context.Background(), conn, input, fn)
}

func describeStacksPagesWithContext(ctx context.Context, conn *appstream.AppStream, input *appstream.DescribeStacksInput, fn func(*appstream.DescribeStacksOutput, bool) bool) error {
	for {
		output, err := conn.DescribeStacksWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func describeUsersPages(conn *appstream.AppStream, input *appstream.DescribeUsersInput, fn func(*appstream.DescribeUsersOutput, bool) bool) error {
	return describeUsersPagesWithContext(context.Background(), conn, input, fn)
}

func describeUsersPagesWithContext(ctx context.Context, conn *appstream.AppStream, input *appstream.DescribeUsersInput, fn func(*appstream.DescribeUsersOutput, bool) bool) error {
	for {
		output, err := conn.DescribeUsersWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func listAssociatedStacksPages(conn *appstream.AppStream, input *appstream.ListAssociatedStacksInput, fn func(*appstream.ListAssociatedStacksOutput, bool) bool) error {
	return listAssociatedStacksPagesWithContext(context.Background(), conn, input, fn)
}

func listAssociatedStacksPagesWithContext(ctx context.Context, conn *appstream.AppStream, input *appstream.ListAssociatedStacksInput, fn func(*appstream.ListAssociatedStacksOutput, bool) bool) error {
	for {
		output, err := conn.ListAssociatedStacksWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}