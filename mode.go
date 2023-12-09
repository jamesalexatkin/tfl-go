package tfl

import (
	"context"
	"fmt"
	"strconv"
)

// GetActiveServiceTypes returns the service type active for a mode. Currently only supports tube.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Mode/Mode_GetActiveServiceTypes
func (c *Client) GetActiveServiceTypes(ctx context.Context) ([]ActiveServiceType, error) {
	path := "/Mode/ActiveServiceTypes"

	activeServiceTypes := []ActiveServiceType{}
	err := c.get(ctx, path, &activeServiceTypes)
	if err != nil {
		return nil, err
	}

	return activeServiceTypes, err
}

// GetArrivalPredictionsForMode gets the next arrival predictions for all stops of a given mode.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Mode/Mode_Arrivals
//
//nolint:revive
func (c *Client) GetArrivalPredictionsForMode(ctx context.Context, mode string, count int) ([]Prediction, error) {
	path := fmt.Sprintf("/Mode/%s/Arrivals", mode)

	predictions := []Prediction{}
	err := c.getWithQueryParams(ctx, path, map[string]string{"count": strconv.Itoa(count)}, &predictions)
	if err != nil {
		return nil, err
	}

	return predictions, err
}
