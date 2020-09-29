package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/wkda/go-libs/net/http/transport"
	"github.com/wkda/go-libs/secret"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]string{"about": "Gogog demo for Auto1!"})
	})

	app.Get("/car/:id/stock-number", func(ctx *fiber.Ctx) error {
		id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		sn, err := fetchStockNumber(id, ctx.Context())
		if err != nil {
			return ctx.SendStatus(fiber.StatusNotFound)
		}

		return ctx.JSON(map[int64]string{id: sn})
	})

	app.Listen(":8080")
}

type CarLead struct {
	ID          int    `json:"id"`
	StockNumber string `json:"stockNumber"`
}

func fetchStockNumber(id int64, ctx context.Context) (string, error) {
	var (
		resp *http.Response
		req  *http.Request
	)

	apiSecret := secret.New(os.Getenv("WA_KEY"), os.Getenv("WA_SECRET"))

	client := &http.Client{
		Transport: transport.Chain(
			transport.NewHeaders(map[string]string{
				"Content-Type": "application/json",
				"Agent":        "golang-all-hands-bdd-demo",
			}),
			transport.NewTraceID(),
			transport.NewSecret(apiSecret),
			transport.NewDebug(true),
		),
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/carlead/%d", os.Getenv("CAR_LEAD_HOST"), id), nil)
	if err != nil {
		return "", err
	}

	req.WithContext(ctx)

	resp, err = client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return "", errors.New("not found")
	} else if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	cl := &CarLead{}
	err = json.Unmarshal(body, cl)
	if err != nil {
		return "", err
	}

	return cl.StockNumber, nil
}
