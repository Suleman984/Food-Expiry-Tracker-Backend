package jobs

import (
	"backend/config"
	"backend/services"
	"fmt"
)

func NotifyUsers(cfg config.Config) {
	items, err := services.GetExpiringItems(cfg)
	if err != nil {
		fmt.Println("Error fetching items:", err)
		return
	}
	//loop started to get emails and send notifications
	if len(items) == 0 {
		fmt.Println("No expiring items found")
		return
	}
	for _, item := range items {
		email, err := services.GetUserEmail(cfg, item.UserID)
		if err != nil {
			fmt.Println("Error fetching email:", err)
			continue
		}

		subject := fmt.Sprintf("Expiry Alert: %s", item.Name)
		body := fmt.Sprintf("Your item '%s' (Qty: %d) is expiring on %s. Please use it before expiry!",
			item.Name, item.Quantity, item.ExpiryDate)

		if err := services.SendEmail(cfg, email, subject, body); err != nil {
			fmt.Println("Error sending email:", err)
		} else {
			fmt.Printf("Email sent to %s for item %s\n", email, item.Name)
		}
	}
}
