package ui

import (
	"context"
	"fmt"
	"strings"

	"github.com/aungshanbo/a9r/aws"
	"github.com/aungshanbo/a9r/models"
	"github.com/aungshanbo/a9r/utils"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ShowEC2Detail(
	app *tview.Application,
	pages *tview.Pages,
	table *tview.Table,
	title string,
	instance *types.Instance,
) {
	if instance == nil {
		return
	}

	text := tview.NewTextView()

	text.SetBorder(true)
	text.SetTitle(" EC2 Detail ")
	text.SetDynamicColors(true)
	text.SetScrollable(true)
	text.SetWordWrap(false)

	var builder strings.Builder

	// ==================================================
	// BASIC INFORMATION
	// ==================================================

	builder.WriteString("[yellow]BASIC INFORMATION[white]\n")
	builder.WriteString(strings.Repeat("─", 70))
	builder.WriteString("\n")

	fmt.Fprintf(
		&builder,
		"Instance ID        : %s\n",
		utils.Stringvalue(instance.InstanceId),
	)

	fmt.Fprintf(
		&builder,
		"Name               : %s\n",
		getEC2Tag(instance.Tags, "Name"),
	)

	fmt.Fprintf(
		&builder,
		"Instance Type      : %s\n",
		string(instance.InstanceType),
	)

	if instance.State != nil {
		fmt.Fprintf(
			&builder,
			"State              : %s\n",
			string(instance.State.Name),
		)

		fmt.Fprintf(
			&builder,
			"State Code         : %d\n",
			utils.Int32value(instance.State.Code),
		)
	} else {
		builder.WriteString("State              : -\n")
		builder.WriteString("State Code         : -\n")
	}

	if instance.StateReason != nil {
		fmt.Fprintf(
			&builder,
			"State Reason       : %s\n",
			utils.Stringvalue(instance.StateReason.Message),
		)
	} else {
		builder.WriteString("State Reason       : -\n")
	}

	fmt.Fprintf(
		&builder,
		"AMI                : %s\n",
		utils.Stringvalue(instance.ImageId),
	)

	fmt.Fprintf(
		&builder,
		"Architecture       : %s\n",
		string(instance.Architecture),
	)

	fmt.Fprintf(
		&builder,
		"Platform           : %s\n",
		string(instance.Platform),
	)

	fmt.Fprintf(
		&builder,
		"Platform Details   : %s\n",
		utils.Stringvalue(instance.PlatformDetails),
	)

	if instance.LaunchTime != nil {
		fmt.Fprintf(
			&builder,
			"Launch Time        : %s\n",
			instance.LaunchTime.Format("2006-01-02 15:04:05 MST"),
		)
	} else {
		builder.WriteString("Launch Time        : -\n")
	}

	// ==================================================
	// COMPUTE
	// ==================================================

	builder.WriteString("\n")
	builder.WriteString("[yellow]COMPUTE[white]\n")
	builder.WriteString(strings.Repeat("─", 70))
	builder.WriteString("\n")

	if instance.CpuOptions != nil {
		fmt.Fprintf(
			&builder,
			"CPU Core Count     : %d\n",
			utils.Int32value(instance.CpuOptions.CoreCount),
		)

		fmt.Fprintf(
			&builder,
			"Threads / Core     : %d\n",
			utils.Int32value(instance.CpuOptions.ThreadsPerCore),
		)
	} else {
		builder.WriteString("CPU Core Count     : -\n")
		builder.WriteString("Threads / Core     : -\n")
	}

	fmt.Fprintf(
		&builder,
		"Hypervisor         : %s\n",
		string(instance.Hypervisor),
	)

	fmt.Fprintf(
		&builder,
		"ENA Support        : %t\n",
		utils.BoolValue(instance.EnaSupport),
	)

	fmt.Fprintf(
		&builder,
		"EBS Optimized      : %t\n",
		utils.BoolValue(instance.EbsOptimized),
	)

	// ==================================================
	// NETWORK
	// ==================================================

	builder.WriteString("\n")
	builder.WriteString("[yellow]NETWORK[white]\n")
	builder.WriteString(strings.Repeat("─", 70))
	builder.WriteString("\n")

	fmt.Fprintf(
		&builder,
		"Private IP         : %s\n",
		utils.Stringvalue(instance.PrivateIpAddress),
	)

	fmt.Fprintf(
		&builder,
		"Public IP          : %s\n",
		utils.Stringvalue(instance.PublicIpAddress),
	)

	fmt.Fprintf(
		&builder,
		"Private DNS        : %s\n",
		utils.Stringvalue(instance.PrivateDnsName),
	)

	fmt.Fprintf(
		&builder,
		"Public DNS         : %s\n",
		utils.Stringvalue(instance.PublicDnsName),
	)

	fmt.Fprintf(
		&builder,
		"VPC ID             : %s\n",
		utils.Stringvalue(instance.VpcId),
	)

	fmt.Fprintf(
		&builder,
		"Subnet ID          : %s\n",
		utils.Stringvalue(instance.SubnetId),
	)

	if instance.Placement != nil {
		fmt.Fprintf(
			&builder,
			"Availability Zone  : %s\n",
			utils.Stringvalue(instance.Placement.AvailabilityZone),
		)

		fmt.Fprintf(
			&builder,
			"AZ ID              : %s\n",
			utils.Stringvalue(instance.Placement.AvailabilityZoneId),
		)

		fmt.Fprintf(
			&builder,
			"Tenancy            : %s\n",
			string(instance.Placement.Tenancy),
		)
	} else {
		builder.WriteString("Availability Zone  : -\n")
		builder.WriteString("AZ ID              : -\n")
		builder.WriteString("Tenancy            : -\n")
	}

	// ==================================================
	// IPV6
	// ==================================================

	builder.WriteString("\n")
	builder.WriteString("[yellow]IPv6 ADDRESSES[white]\n")
	builder.WriteString(strings.Repeat("─", 70))
	builder.WriteString("\n")

	if instance.Ipv6Address == nil {
		builder.WriteString("  -\n")
	} else {
		fmt.Fprintf(
			&builder,
			"  %s\n",
			utils.Stringvalue(instance.Ipv6Address),
		)
	}

	// ==================================================
	// SECURITY
	// ==================================================

	builder.WriteString("\n")
	builder.WriteString("[yellow]SECURITY[white]\n")
	builder.WriteString(strings.Repeat("─", 70))
	builder.WriteString("\n")

	fmt.Fprintf(
		&builder,
		"Key Name           : %s\n",
		utils.Stringvalue(instance.KeyName),
	)

	if instance.IamInstanceProfile != nil {
		fmt.Fprintf(
			&builder,
			"IAM Profile        : %s\n",
			utils.Stringvalue(instance.IamInstanceProfile.Arn),
		)
	} else {
		builder.WriteString("IAM Profile        : -\n")
	}

	// ==================================================
	// SECURITY GROUPS
	// ==================================================

	builder.WriteString("\n")
	builder.WriteString("[yellow]SECURITY GROUPS[white]\n")
	builder.WriteString(strings.Repeat("─", 70))
	builder.WriteString("\n")

	if len(instance.SecurityGroups) == 0 {
		builder.WriteString("  -\n")
	} else {
		for _, sg := range instance.SecurityGroups {
			fmt.Fprintf(
				&builder,
				"  %-25s %s\n",
				utils.Stringvalue(sg.GroupName),
				utils.Stringvalue(sg.GroupId),
			)
		}
	}

	// ==================================================
	// NETWORK INTERFACES
	// ==================================================

	builder.WriteString("\n")
	builder.WriteString("[yellow]NETWORK INTERFACES[white]\n")
	builder.WriteString(strings.Repeat("─", 70))
	builder.WriteString("\n")

	if len(instance.NetworkInterfaces) == 0 {
		builder.WriteString("  -\n")
	} else {

		for i, ni := range instance.NetworkInterfaces {

			fmt.Fprintf(
				&builder,
				"\n  Interface #%d\n",
				i+1,
			)

			fmt.Fprintf(
				&builder,
				"    ID             : %s\n",
				utils.Stringvalue(ni.NetworkInterfaceId),
			)

			fmt.Fprintf(
				&builder,
				"    Description    : %s\n",
				utils.Stringvalue(ni.Description),
			)

			fmt.Fprintf(
				&builder,
				"    Status         : %s\n",
				string(ni.Status),
			)

			fmt.Fprintf(
				&builder,
				"    MAC Address    : %s\n",
				utils.Stringvalue(ni.MacAddress),
			)

			fmt.Fprintf(
				&builder,
				"    Private DNS    : %s\n",
				utils.Stringvalue(ni.PrivateDnsName),
			)

			fmt.Fprintf(
				&builder,
				"    VPC ID         : %s\n",
				utils.Stringvalue(ni.VpcId),
			)

			fmt.Fprintf(
				&builder,
				"    Subnet ID      : %s\n",
				utils.Stringvalue(ni.SubnetId),
			)

			if len(ni.PrivateIpAddresses) > 0 {
				builder.WriteString("    Private IPs     :\n")

				for _, ip := range ni.PrivateIpAddresses {
					fmt.Fprintf(
						&builder,
						"      - %s\n",
						utils.Stringvalue(ip.PrivateIpAddress),
					)
				}
			}
		}
	}

	// ==================================================
	// STORAGE
	// ==================================================

	builder.WriteString("\n")
	builder.WriteString("[yellow]STORAGE[white]\n")
	builder.WriteString(strings.Repeat("─", 70))
	builder.WriteString("\n")

	if instance.RootDeviceName != nil {
		fmt.Fprintf(
			&builder,
			"Root Device       : %s\n",
			utils.Stringvalue(instance.RootDeviceName),
		)
	} else {
		builder.WriteString("Root Device       : -\n")
	}

	fmt.Fprintf(
		&builder,
		"Root Device Type  : %s\n",
		string(instance.RootDeviceType),
	)

	if len(instance.BlockDeviceMappings) == 0 {
		builder.WriteString("\nBlock Devices     : -\n")
	} else {

		builder.WriteString("\nBlock Devices:\n")

		for i, device := range instance.BlockDeviceMappings {

			fmt.Fprintf(
				&builder,
				"\n  Device #%d\n",
				i+1,
			)

			fmt.Fprintf(
				&builder,
				"    Device Name    : %s\n",
				utils.Stringvalue(device.DeviceName),
			)

			if device.Ebs != nil {

				fmt.Fprintf(
					&builder,
					"    Volume ID      : %s\n",
					utils.Stringvalue(device.Ebs.VolumeId),
				)

				fmt.Fprintf(
					&builder,
					"    Delete on Term.: %t\n",
					utils.BoolValue(device.Ebs.DeleteOnTermination),
				)

			}
		}
	}

	// ==================================================
	// PLACEMENT
	// ==================================================

	builder.WriteString("\n")
	builder.WriteString("[yellow]PLACEMENT[white]\n")
	builder.WriteString(strings.Repeat("─", 70))
	builder.WriteString("\n")

	if instance.Placement != nil {

		fmt.Fprintf(
			&builder,
			"Availability Zone : %s\n",
			utils.Stringvalue(instance.Placement.AvailabilityZone),
		)

		fmt.Fprintf(
			&builder,
			"AZ ID             : %s\n",
			utils.Stringvalue(instance.Placement.AvailabilityZoneId),
		)

		fmt.Fprintf(
			&builder,
			"Tenancy           : %s\n",
			string(instance.Placement.Tenancy),
		)

		fmt.Fprintf(
			&builder,
			"Host ID           : %s\n",
			utils.Stringvalue(instance.Placement.HostId),
		)

		fmt.Fprintf(
			&builder,
			"Partition Number  : %d\n",
			utils.Int32value(instance.Placement.PartitionNumber),
		)

		if instance.Placement.GroupName != nil {
			fmt.Fprintf(
				&builder,
				"Placement Group   : %s\n",
				utils.Stringvalue(instance.Placement.GroupName),
			)
		}
	}

	// ==================================================
	// MONITORING
	// ==================================================

	builder.WriteString("\n")
	builder.WriteString("[yellow]MONITORING[white]\n")
	builder.WriteString(strings.Repeat("─", 70))
	builder.WriteString("\n")

	fmt.Fprintf(
		&builder,
		"Monitoring        : %s\n",
		string(instance.Monitoring.State),
	)

	// ==================================================
	// LIFECYCLE
	// ==================================================

	builder.WriteString("\n")
	builder.WriteString("[yellow]LIFECYCLE[white]\n")
	builder.WriteString(strings.Repeat("─", 70))
	builder.WriteString("\n")

	fmt.Fprintf(
		&builder,
		"Lifecycle         : %s\n",
		string(instance.InstanceLifecycle),
	)

	fmt.Fprintf(
		&builder,
		"Spot Request ID   : %s\n",
		utils.Stringvalue(instance.SpotInstanceRequestId),
	)

	// ==================================================
	// METADATA
	// ==================================================

	builder.WriteString("\n")
	builder.WriteString("[yellow]INSTANCE METADATA[white]\n")
	builder.WriteString(strings.Repeat("─", 70))
	builder.WriteString("\n")

	if instance.MetadataOptions != nil {

		fmt.Fprintf(
			&builder,
			"HTTP Endpoint     : %s\n",
			string(instance.MetadataOptions.HttpEndpoint),
		)

		fmt.Fprintf(
			&builder,
			"HTTP Tokens       : %s\n",
			string(instance.MetadataOptions.HttpTokens),
		)

		fmt.Fprintf(
			&builder,
			"Hop Limit         : %d\n",
			utils.Int32value(instance.MetadataOptions.HttpPutResponseHopLimit),
		)
	}

	// ==================================================
	// TAGS
	// ==================================================

	builder.WriteString("\n")
	builder.WriteString("[yellow]TAGS[white]\n")
	builder.WriteString(strings.Repeat("─", 70))
	builder.WriteString("\n")

	if len(instance.Tags) == 0 {
		builder.WriteString("  -\n")
	} else {

		for _, tag := range instance.Tags {

			fmt.Fprintf(
				&builder,
				"  %-25s : %s\n",
				utils.Stringvalue(tag.Key),
				utils.Stringvalue(tag.Value),
			)
		}
	}

	// ==================================================
	// DISPLAY
	// ==================================================

	text.SetText(builder.String())

	// ==================================================
	// KEY HANDLING
	// ==================================================

	text.SetInputCapture(
		func(event *tcell.EventKey) *tcell.EventKey {

			switch event.Key() {

			case tcell.KeyEscape:

				pages.RemovePage("ec2-detail")
				app.SetFocus(table)

				return nil
			}

			return event
		},
	)

	// ==================================================
	// SHOW
	// ==================================================

	pages.AddPage(
		"ec2-detail",
		text,
		true,
		true,
	)

	app.SetFocus(text)
}

// ======================================================
// GET EC2 TAG
// ======================================================

func getEC2Tag(
	tags []types.Tag,
	key string,
) string {

	for _, tag := range tags {

		if utils.Stringvalue(tag.Key) == key {
			return utils.Stringvalue(tag.Value)
		}
	}

	return "-"
}

func ShowS3Detail(
	app *tview.Application,
	pages *tview.Pages,
	table *tview.Table,
	profile string,
	detail *models.S3BucketDetail,
) {

	text := tview.NewTextView()

	text.SetDynamicColors(true)
	text.SetBorder(true)
	text.SetTitle(" S3 Bucket Detail ")
	text.SetScrollable(true)

	detail.ObjectCount = -1
	detail.SizeBytes = -1

	render := func() {

		var b strings.Builder

		b.WriteString("\n")

		writeRow := func(label, value string) {
			b.WriteString(
				fmt.Sprintf(
					"  %-20s %s\n",
					label,
					value,
				),
			)
		}

		writeRow("Name", detail.Name)
		writeRow("Region", detail.Region)

		created := "-"

		if !detail.CreationDate.IsZero() {
			created = detail.CreationDate.Format("2006-01-02")
		}

		writeRow("Created", created)

		b.WriteString("\n")

		objectCount := "Loading..."

		if detail.ObjectCount >= 0 {
			objectCount = fmt.Sprintf(
				"%d",
				detail.ObjectCount,
			)
		}

		size := "Loading..."

		if detail.SizeBytes >= 0 {
			size = aws.FormatBytes(detail.SizeBytes)
		}

		writeRow("Objects", objectCount)
		writeRow("Total Size", size)

		b.WriteString("\n")

		writeRow("Versioning", detail.Versioning)
		writeRow("Encryption", detail.Encryption)
		writeRow("Object Lock", detail.ObjectLock)

		b.WriteString("\n")

		writeRow("Public Access", detail.PublicAccess)
		writeRow("Object Ownership", detail.ObjectOwnership)
		writeRow("ACL", detail.ACL)

		b.WriteString("\n")

		writeRow("Policy", detail.Policy)

		writeRow(
			"Lifecycle Rules",
			fmt.Sprintf("%d", detail.LifecycleRules),
		)

		writeRow("Replication", detail.Replication)
		writeRow("Access Logging", detail.AccessLogging)

		b.WriteString("\n")

		writeRow(
			"Tags",
			fmt.Sprintf("%d", len(detail.Tags)),
		)

		b.WriteString("\n")
		b.WriteString("  ESC / q = Close\n")

		text.SetText(b.String())
	}

	render()

	text.SetInputCapture(
		func(event *tcell.EventKey) *tcell.EventKey {

			if event.Key() == tcell.KeyEsc ||
				event.Rune() == 'q' {

				pages.RemovePage("s3-detail")
				app.SetFocus(table)

				return nil
			}

			return event
		},
	)

	pages.AddPage(
		"s3-detail",
		Center(100, 30, text),
		true,
		true,
	)

	app.SetFocus(text)

	// ==========================================
	// BACKGROUND STATISTICS
	// ==========================================

	go func() {

		objectCount, sizeBytes :=
			aws.GetS3BucketStatistics(
				context.Background(),
				profile,
				detail.Region,
				detail.Name,
			)

		app.QueueUpdateDraw(
			func() {

				if !pages.HasPage("s3-detail") {
					return
				}

				detail.ObjectCount = objectCount
				detail.SizeBytes = sizeBytes

				render()
			},
		)
	}()
}
