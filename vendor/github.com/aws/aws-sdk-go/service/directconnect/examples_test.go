// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package directconnect_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/directconnect"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleDirectConnect_AllocateConnectionOnInterconnect() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.AllocateConnectionOnInterconnectInput{
		Bandwidth:      aws.String("Bandwidth"),      // Required
		ConnectionName: aws.String("ConnectionName"), // Required
		InterconnectId: aws.String("InterconnectId"), // Required
		OwnerAccount:   aws.String("OwnerAccount"),   // Required
		Vlan:           aws.Int64(1),                 // Required
	}
	resp, err := svc.AllocateConnectionOnInterconnect(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_AllocateHostedConnection() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.AllocateHostedConnectionInput{
		Bandwidth:      aws.String("Bandwidth"),      // Required
		ConnectionId:   aws.String("ConnectionId"),   // Required
		ConnectionName: aws.String("ConnectionName"), // Required
		OwnerAccount:   aws.String("OwnerAccount"),   // Required
		Vlan:           aws.Int64(1),                 // Required
	}
	resp, err := svc.AllocateHostedConnection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_AllocatePrivateVirtualInterface() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.AllocatePrivateVirtualInterfaceInput{
		ConnectionId: aws.String("ConnectionId"), // Required
		NewPrivateVirtualInterfaceAllocation: &directconnect.NewPrivateVirtualInterfaceAllocation{ // Required
			Asn:                  aws.Int64(1),                       // Required
			VirtualInterfaceName: aws.String("VirtualInterfaceName"), // Required
			Vlan:                 aws.Int64(1),                       // Required
			AddressFamily:        aws.String("AddressFamily"),
			AmazonAddress:        aws.String("AmazonAddress"),
			AuthKey:              aws.String("BGPAuthKey"),
			CustomerAddress:      aws.String("CustomerAddress"),
		},
		OwnerAccount: aws.String("OwnerAccount"), // Required
	}
	resp, err := svc.AllocatePrivateVirtualInterface(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_AllocatePublicVirtualInterface() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.AllocatePublicVirtualInterfaceInput{
		ConnectionId: aws.String("ConnectionId"), // Required
		NewPublicVirtualInterfaceAllocation: &directconnect.NewPublicVirtualInterfaceAllocation{ // Required
			Asn:                  aws.Int64(1),                       // Required
			VirtualInterfaceName: aws.String("VirtualInterfaceName"), // Required
			Vlan:                 aws.Int64(1),                       // Required
			AddressFamily:        aws.String("AddressFamily"),
			AmazonAddress:        aws.String("AmazonAddress"),
			AuthKey:              aws.String("BGPAuthKey"),
			CustomerAddress:      aws.String("CustomerAddress"),
			RouteFilterPrefixes: []*directconnect.RouteFilterPrefix{
				{ // Required
					Cidr: aws.String("CIDR"),
				},
				// More values...
			},
		},
		OwnerAccount: aws.String("OwnerAccount"), // Required
	}
	resp, err := svc.AllocatePublicVirtualInterface(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_AssociateConnectionWithLag() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.AssociateConnectionWithLagInput{
		ConnectionId: aws.String("ConnectionId"), // Required
		LagId:        aws.String("LagId"),        // Required
	}
	resp, err := svc.AssociateConnectionWithLag(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_AssociateHostedConnection() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.AssociateHostedConnectionInput{
		ConnectionId:       aws.String("ConnectionId"), // Required
		ParentConnectionId: aws.String("ConnectionId"), // Required
	}
	resp, err := svc.AssociateHostedConnection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_AssociateVirtualInterface() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.AssociateVirtualInterfaceInput{
		ConnectionId:       aws.String("ConnectionId"),       // Required
		VirtualInterfaceId: aws.String("VirtualInterfaceId"), // Required
	}
	resp, err := svc.AssociateVirtualInterface(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_ConfirmConnection() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.ConfirmConnectionInput{
		ConnectionId: aws.String("ConnectionId"), // Required
	}
	resp, err := svc.ConfirmConnection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_ConfirmPrivateVirtualInterface() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.ConfirmPrivateVirtualInterfaceInput{
		VirtualGatewayId:   aws.String("VirtualGatewayId"),   // Required
		VirtualInterfaceId: aws.String("VirtualInterfaceId"), // Required
	}
	resp, err := svc.ConfirmPrivateVirtualInterface(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_ConfirmPublicVirtualInterface() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.ConfirmPublicVirtualInterfaceInput{
		VirtualInterfaceId: aws.String("VirtualInterfaceId"), // Required
	}
	resp, err := svc.ConfirmPublicVirtualInterface(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_CreateBGPPeer() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.CreateBGPPeerInput{
		NewBGPPeer: &directconnect.NewBGPPeer{
			AddressFamily:   aws.String("AddressFamily"),
			AmazonAddress:   aws.String("AmazonAddress"),
			Asn:             aws.Int64(1),
			AuthKey:         aws.String("BGPAuthKey"),
			CustomerAddress: aws.String("CustomerAddress"),
		},
		VirtualInterfaceId: aws.String("VirtualInterfaceId"),
	}
	resp, err := svc.CreateBGPPeer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_CreateConnection() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.CreateConnectionInput{
		Bandwidth:      aws.String("Bandwidth"),      // Required
		ConnectionName: aws.String("ConnectionName"), // Required
		Location:       aws.String("LocationCode"),   // Required
		LagId:          aws.String("LagId"),
	}
	resp, err := svc.CreateConnection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_CreateInterconnect() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.CreateInterconnectInput{
		Bandwidth:        aws.String("Bandwidth"),        // Required
		InterconnectName: aws.String("InterconnectName"), // Required
		Location:         aws.String("LocationCode"),     // Required
		LagId:            aws.String("LagId"),
	}
	resp, err := svc.CreateInterconnect(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_CreateLag() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.CreateLagInput{
		ConnectionsBandwidth: aws.String("Bandwidth"),    // Required
		LagName:              aws.String("LagName"),      // Required
		Location:             aws.String("LocationCode"), // Required
		NumberOfConnections:  aws.Int64(1),               // Required
		ConnectionId:         aws.String("ConnectionId"),
	}
	resp, err := svc.CreateLag(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_CreatePrivateVirtualInterface() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.CreatePrivateVirtualInterfaceInput{
		ConnectionId: aws.String("ConnectionId"), // Required
		NewPrivateVirtualInterface: &directconnect.NewPrivateVirtualInterface{ // Required
			Asn:                  aws.Int64(1),                       // Required
			VirtualGatewayId:     aws.String("VirtualGatewayId"),     // Required
			VirtualInterfaceName: aws.String("VirtualInterfaceName"), // Required
			Vlan:                 aws.Int64(1),                       // Required
			AddressFamily:        aws.String("AddressFamily"),
			AmazonAddress:        aws.String("AmazonAddress"),
			AuthKey:              aws.String("BGPAuthKey"),
			CustomerAddress:      aws.String("CustomerAddress"),
		},
	}
	resp, err := svc.CreatePrivateVirtualInterface(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_CreatePublicVirtualInterface() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.CreatePublicVirtualInterfaceInput{
		ConnectionId: aws.String("ConnectionId"), // Required
		NewPublicVirtualInterface: &directconnect.NewPublicVirtualInterface{ // Required
			Asn:                  aws.Int64(1),                       // Required
			VirtualInterfaceName: aws.String("VirtualInterfaceName"), // Required
			Vlan:                 aws.Int64(1),                       // Required
			AddressFamily:        aws.String("AddressFamily"),
			AmazonAddress:        aws.String("AmazonAddress"),
			AuthKey:              aws.String("BGPAuthKey"),
			CustomerAddress:      aws.String("CustomerAddress"),
			RouteFilterPrefixes: []*directconnect.RouteFilterPrefix{
				{ // Required
					Cidr: aws.String("CIDR"),
				},
				// More values...
			},
		},
	}
	resp, err := svc.CreatePublicVirtualInterface(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DeleteBGPPeer() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DeleteBGPPeerInput{
		Asn:                aws.Int64(1),
		CustomerAddress:    aws.String("CustomerAddress"),
		VirtualInterfaceId: aws.String("VirtualInterfaceId"),
	}
	resp, err := svc.DeleteBGPPeer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DeleteConnection() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DeleteConnectionInput{
		ConnectionId: aws.String("ConnectionId"), // Required
	}
	resp, err := svc.DeleteConnection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DeleteInterconnect() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DeleteInterconnectInput{
		InterconnectId: aws.String("InterconnectId"), // Required
	}
	resp, err := svc.DeleteInterconnect(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DeleteLag() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DeleteLagInput{
		LagId: aws.String("LagId"), // Required
	}
	resp, err := svc.DeleteLag(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DeleteVirtualInterface() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DeleteVirtualInterfaceInput{
		VirtualInterfaceId: aws.String("VirtualInterfaceId"), // Required
	}
	resp, err := svc.DeleteVirtualInterface(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DescribeConnectionLoa() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DescribeConnectionLoaInput{
		ConnectionId:   aws.String("ConnectionId"), // Required
		LoaContentType: aws.String("LoaContentType"),
		ProviderName:   aws.String("ProviderName"),
	}
	resp, err := svc.DescribeConnectionLoa(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DescribeConnections() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DescribeConnectionsInput{
		ConnectionId: aws.String("ConnectionId"),
	}
	resp, err := svc.DescribeConnections(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DescribeConnectionsOnInterconnect() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DescribeConnectionsOnInterconnectInput{
		InterconnectId: aws.String("InterconnectId"), // Required
	}
	resp, err := svc.DescribeConnectionsOnInterconnect(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DescribeHostedConnections() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DescribeHostedConnectionsInput{
		ConnectionId: aws.String("ConnectionId"), // Required
	}
	resp, err := svc.DescribeHostedConnections(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DescribeInterconnectLoa() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DescribeInterconnectLoaInput{
		InterconnectId: aws.String("InterconnectId"), // Required
		LoaContentType: aws.String("LoaContentType"),
		ProviderName:   aws.String("ProviderName"),
	}
	resp, err := svc.DescribeInterconnectLoa(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DescribeInterconnects() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DescribeInterconnectsInput{
		InterconnectId: aws.String("InterconnectId"),
	}
	resp, err := svc.DescribeInterconnects(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DescribeLags() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DescribeLagsInput{
		LagId: aws.String("LagId"),
	}
	resp, err := svc.DescribeLags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DescribeLoa() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DescribeLoaInput{
		ConnectionId:   aws.String("ConnectionId"), // Required
		LoaContentType: aws.String("LoaContentType"),
		ProviderName:   aws.String("ProviderName"),
	}
	resp, err := svc.DescribeLoa(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DescribeLocations() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	var params *directconnect.DescribeLocationsInput
	resp, err := svc.DescribeLocations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DescribeTags() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DescribeTagsInput{
		ResourceArns: []*string{ // Required
			aws.String("ResourceArn"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DescribeVirtualGateways() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	var params *directconnect.DescribeVirtualGatewaysInput
	resp, err := svc.DescribeVirtualGateways(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DescribeVirtualInterfaces() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DescribeVirtualInterfacesInput{
		ConnectionId:       aws.String("ConnectionId"),
		VirtualInterfaceId: aws.String("VirtualInterfaceId"),
	}
	resp, err := svc.DescribeVirtualInterfaces(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_DisassociateConnectionFromLag() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.DisassociateConnectionFromLagInput{
		ConnectionId: aws.String("ConnectionId"), // Required
		LagId:        aws.String("LagId"),        // Required
	}
	resp, err := svc.DisassociateConnectionFromLag(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_TagResource() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.TagResourceInput{
		ResourceArn: aws.String("ResourceArn"), // Required
		Tags: []*directconnect.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKey"), // Required
				Value: aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.TagResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_UntagResource() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.UntagResourceInput{
		ResourceArn: aws.String("ResourceArn"), // Required
		TagKeys: []*string{ // Required
			aws.String("TagKey"), // Required
			// More values...
		},
	}
	resp, err := svc.UntagResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDirectConnect_UpdateLag() {
	sess := session.Must(session.NewSession())

	svc := directconnect.New(sess)

	params := &directconnect.UpdateLagInput{
		LagId:        aws.String("LagId"), // Required
		LagName:      aws.String("LagName"),
		MinimumLinks: aws.Int64(1),
	}
	resp, err := svc.UpdateLag(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
