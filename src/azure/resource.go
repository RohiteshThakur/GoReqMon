package azure

import (
	"sync"
	"fmt"
	"context"
	"strings"
	"log"
	//"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"reflect"
)

func GetResourceGroups(sess *AzureSession) ([]string, error) {
	tab := make([]string, 0)
	var err error

	grClient := resources.NewGroupsClient(sess.SubscriptionId)
	grClient.Authorizer = sess.Authorizer

	for list, err := grClient.ListComplete(context.Background(), "", nil); list.NotDone(); err = list.Next() {
		if err != nil {
			return nil, err
		}
		rgName := *list.Value().Name
		fmt.Println(rgName)
		tab = append(tab, rgName)
	}
	return tab, err
}

func GetVM(sess *AzureSession, rg string, wg *sync.WaitGroup) {
	defer wg.Done()

	vmClient := compute.NewVirtualMachinesClient(sess.SubscriptionId)
	vmClient.Authorizer = sess.Authorizer

	for vm, err := vmClient.ListComplete(context.Background(), rg); vm.NotDone(); err = vm.Next() {
		if err != nil {
			log.Print("got error while traverising RG list: ", err)
		}

		i := vm.Value()
		fmt.Printf("\n\n%s,%s,%s\n", rg, *i.Name, *i.ID)

		tags := []string{}
		fmt.Println("---------------------------")
		for k, v := range i.Tags {
			//fmt.Println("Tags: ", k, *v)
			tags = append(tags, fmt.Sprintf("%s?%s", k, *v))
		}
		tagsS := strings.Join(tags, "%")

		if len(i.Tags) > 0 {
			//fmt.Printf("%s,%s,%s,<%s>\n", rg, *i.Name, *i.ID, tagsS)
			fmt.Printf("%s, <%s>\\n",*i.Name, tagsS)
		} else {
			fmt.Printf("%s,%s,%s\n", rg, *i.Name, *i.ID)
		}
	}
}

func GetADGroups(sess *AzureSession) (graphrbac.GroupListResultPage, error) {

	groupsClient := graphrbac.NewGroupsClient(sess.TenantId)
	groupsClient.Authorizer = sess.Authorizer

	adGroupNames, err := groupsClient.List(context.Background(), "")

	fmt.Println(reflect.TypeOf(adGroupNames))
	objectType := reflect.TypeOf(adGroupNames)
	for i := 0; i < objectType.NumMethod(); i++ {
		method := objectType.Method(i)
		fmt.Println(method.Name)
	}	

	fmt.Println(adGroupNames.Values())
	return adGroupNames, err
}


