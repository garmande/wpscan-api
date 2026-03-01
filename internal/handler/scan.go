package handler

import (
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/garmande/wpscan-api/internal/model"
	"github.com/garmande/wpscan-api/internal/scanner"
	"github.com/labstack/echo/v4"
)

// In-memory store — replace with a database later.
var (
	scans   = map[string]*model.Scan{}
	scansMu sync.RWMutex
)

type createScanRequest struct {
	URL string `json:"url" validate:"required"`
}

func CreateScan(c echo.Context) error {
	var req createScanRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}
	if req.URL == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "url is required"})
	}

	scan := &model.Scan{
		ID:        uuid.NewString(),
		URL:       req.URL,
		Status:    model.ScanStatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	scansMu.Lock()
	scans[scan.ID] = scan
	scansMu.Unlock()

	// Run the scan asynchronously.
	go func() {
		scansMu.Lock()
		scan.Status = model.ScanStatusRunning
		scan.UpdatedAt = time.Now()
		scansMu.Unlock()

		results, err := scanner.Run(scan.URL)

		scansMu.Lock()
		defer scansMu.Unlock()
		scan.UpdatedAt = time.Now()
		if err != nil {
			scan.Status = model.ScanStatusFailed
			scan.Error = err.Error()
		} else {
			scan.Status = model.ScanStatusDone
			scan.Results = results
		}
	}()

	return c.JSON(http.StatusAccepted, scan)
}

func ListScans(c echo.Context) error {
	scansMu.RLock()
	defer scansMu.RUnlock()

	list := make([]*model.Scan, 0, len(scans))
	for _, s := range scans {
		list = append(list, s)
	}
	return c.JSON(http.StatusOK, list)
}

func GetScan(c echo.Context) error {
	id := c.Param("id")

	scansMu.RLock()
	s, ok := scans[id]
	scansMu.RUnlock()

	if !ok {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "scan not found"})
	}
	return c.JSON(http.StatusOK, s)
}
