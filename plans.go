// plans.go
package main

// https://github.com/stripe/stripe-go/blob/master/v32_migration_guide.md

import (
	"flag"
	"fmt"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go/product"
)

func main() {
	var createPlans bool
	var deletePlans bool
	var listPlans bool
	var deletePlan string
	var deleteProduct string

	// Set your secret key: remember to change this to your live secret key in production
	// See your keys here: https://dashboard.stripe.com/account/apikeys
	stripe.Key = os.Getenv("STRIPE_PRIVATE_KEY")

	// flags

	flag.BoolVar(&createPlans, "create_plans", false, "Create product and all plans")
	flag.BoolVar(&deletePlans, "delete_plans", false, "Delete all plans")
	flag.BoolVar(&listPlans, "list", false, "List all plans")
	flag.StringVar(&deleteProduct, "delprod", "", "Delete product")
	flag.StringVar(&deletePlan, "delplan", "", "Delete a plan given an ID")

	flag.Parse()

	if createPlans == true {
		// Create all plans

		paramsProduct := &stripe.ProductParams{
			Name: stripe.String("Example Service"),
			Type: stripe.String(string(stripe.ProductTypeService)),
		}
		prod, errProduct := product.New(paramsProduct)

		if errProduct != nil {
			fmt.Println(errProduct)
		} else {
			fmt.Println(prod)
		}

		// Bronze Monthly

		paramsPlan := &stripe.PlanParams{
			Nickname:      stripe.String("Bronze Volume Pricing"),
			ProductID:     stripe.String(prod.ID),
			Currency:      stripe.String(string(stripe.CurrencyUSD)),
			Interval:      stripe.String(string(stripe.PlanIntervalMonth)),
			UsageType:     stripe.String(string(stripe.PlanUsageTypeMetered)),
			BillingScheme: stripe.String(string(stripe.PlanBillingSchemeTiered)),
			TiersMode:     stripe.String(string(stripe.PlanTiersModeVolume)),
			Tiers: []*stripe.PlanTierParams{
				{
					FlatAmount: stripe.Int64(900),
					UnitAmount: stripe.Int64(20),
					UpTo:       stripe.Int64(5),
				}, {
					FlatAmount: stripe.Int64(1000),
					UnitAmount: stripe.Int64(15),
					UpTo:       stripe.Int64(100),
				}, {
					FlatAmount: stripe.Int64(1500),
					UnitAmount: stripe.Int64(10),
					UpTo:       stripe.Int64(500),
				}, {
					FlatAmount: stripe.Int64(2000),
					UnitAmount: stripe.Int64(5),
					UpToInf:    stripe.Bool(true),
				},
			},
		}

		// TODO: Change roles of each plan to role of user for each plan
		paramsPlan.AddMetadata("plankey", "bronze")
		paramsPlan.AddMetadata("role", "bronze")

		p, errPlan := plan.New(paramsPlan)

		if errPlan != nil {
			fmt.Println(errPlan)
		} else {
			fmt.Println(p)
		}

		// Silver Monthly

		paramsPlan = &stripe.PlanParams{
			Nickname:      stripe.String("Silver Volume Pricing"),
			ProductID:     stripe.String(prod.ID),
			Currency:      stripe.String(string(stripe.CurrencyUSD)),
			Interval:      stripe.String(string(stripe.PlanIntervalMonth)),
			UsageType:     stripe.String(string(stripe.PlanUsageTypeMetered)),
			BillingScheme: stripe.String(string(stripe.PlanBillingSchemeTiered)),
			TiersMode:     stripe.String(string(stripe.PlanTiersModeVolume)),
			Tiers: []*stripe.PlanTierParams{
				{
					FlatAmount: stripe.Int64(900),
					UnitAmount: stripe.Int64(20),
					UpTo:       stripe.Int64(5),
				}, {
					FlatAmount: stripe.Int64(1000),
					UnitAmount: stripe.Int64(15),
					UpTo:       stripe.Int64(100),
				}, {
					FlatAmount: stripe.Int64(1500),
					UnitAmount: stripe.Int64(10),
					UpTo:       stripe.Int64(500),
				}, {
					FlatAmount: stripe.Int64(2000),
					UnitAmount: stripe.Int64(5),
					UpToInf:    stripe.Bool(true),
				},
			},
		}

		paramsPlan.AddMetadata("plankey", "silver")
		paramsPlan.AddMetadata("role", "silver")

		p, errPlan = plan.New(paramsPlan)

		if errPlan != nil {
			fmt.Println(errPlan)
		} else {
			fmt.Println(p)
		}

		// Gold Monthly

		paramsPlan = &stripe.PlanParams{
			Nickname:      stripe.String("Gold Volume Pricing"),
			ProductID:     stripe.String(prod.ID),
			Currency:      stripe.String(string(stripe.CurrencyUSD)),
			Interval:      stripe.String(string(stripe.PlanIntervalMonth)),
			UsageType:     stripe.String(string(stripe.PlanUsageTypeMetered)),
			BillingScheme: stripe.String(string(stripe.PlanBillingSchemeTiered)),
			TiersMode:     stripe.String(string(stripe.PlanTiersModeVolume)),
			Tiers: []*stripe.PlanTierParams{
				{
					FlatAmount: stripe.Int64(900),
					UnitAmount: stripe.Int64(20),
					UpTo:       stripe.Int64(5),
				}, {
					FlatAmount: stripe.Int64(1000),
					UnitAmount: stripe.Int64(15),
					UpTo:       stripe.Int64(100),
				}, {
					FlatAmount: stripe.Int64(1500),
					UnitAmount: stripe.Int64(10),
					UpTo:       stripe.Int64(500),
				}, {
					FlatAmount: stripe.Int64(2000),
					UnitAmount: stripe.Int64(5),
					UpToInf:    stripe.Bool(true),
				},
			},
		}

		paramsPlan.AddMetadata("plankey", "gold")
		paramsPlan.AddMetadata("role", "gold")

		p, errPlan = plan.New(paramsPlan)

		if errPlan != nil {
			fmt.Println(errPlan)
		} else {
			fmt.Println(p)
		}
	} else if deletePlans == true {
		// Delete all plans

		plans := make([]string, 0, 4)

		// Get list of plans

		params := &stripe.PlanListParams{}
		params.Filters.AddFilter("limit", "", "9")
		i := plan.List(params)
		for i.Next() {
			p := i.Plan()
			plans = append(plans, p.ID)
		}

		// Delete each plan

		for _, id := range plans {
			p, err := plan.Del(id, nil)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(p)
			}
		}
	} else if listPlans == true {
		// List all plans

		params := &stripe.PlanListParams{}
		params.Filters.AddFilter("limit", "", "9")
		i := plan.List(params)
		for i.Next() {
			p := i.Plan()
			fmt.Println(p)
		}

		i = plan.List(params)
		fmt.Println("{")
		for i.Next() {
			p := i.Plan()
			fmt.Printf("    \"%s\": \"%s\",\n", p.Metadata["plankey"], p.ID)
		}
		fmt.Println("}")
	} else if deletePlan != "" {
		// Delete a plan given an ID

		p, err := plan.Del(deletePlan, nil)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(p)
		}
	} else if deleteProduct != "" {
		// Delete a product given an ID

		p, err := product.Del(deleteProduct, nil)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(p)
		}
	}
}
