package geverse

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

func (g *Geverse) GetNotices(ctx context.Context, communityID int64) ([]Notice, error) {
	req, err := g.newBasicRequest(ctx, http.MethodGet, endpointNotices(communityID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received unexpected status code: expected 200, received %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var notices []Notice

	doc.Find(".content__item").Each(func(_ int, selection *goquery.Selection) {
		var notice Notice

		noticeLink, _ := selection.Find("a").Attr("href")
		if noticeLink == "" {
			return
		}

		noticeLinkParts := strings.Split(noticeLink, "/")
		noticeLinkPartID := noticeLinkParts[len(noticeLinkParts)-1]

		notice.ID, err = strconv.Atoi(noticeLinkPartID)
		if err != nil {
			err = errors.Wrapf(err, "unexpected notice ID: %s", noticeLinkPartID)
			return
		}

		notice.Link = noticeLink
		if notice.Link != "" && strings.HasPrefix(notice.Link, "/") {
			notice.Link = strings.TrimRight(endpointBaseURL, "/") + notice.Link
		}

		notice.Title = selection.Find(".item__title").Text()
		notice.Label = selection.Find(".item__label").Text()

		dateText := selection.Find(".item__date").Text()
		if dateText != "" {
			dateParsed, err := time.Parse("2006.01.02", dateText)
			if err != nil {
				err = errors.Wrapf(err, "date in unexpected format: %s", dateText)
				return
			}

			notice.Date = dateParsed
		}

		if notice.ID == 0 ||
			notice.Link == "" ||
			notice.Title == "" {
			return
		}

		notices = append(notices, notice)
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse notice")
	}

	if len(notices) <= 0 {
		return nil, errors.New("no notices found")
	}

	return notices, nil
}
