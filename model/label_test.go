package model

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/ngavinsir/clickbait/models"
)

func TestLabel(t *testing.T) {
	db, err := ConnectTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		ResetTestDB(db)
		db.Close()
	}()
	ResetTestDB(db)

	testRepository := initTestRepository(db)

	user, err := testRepository.CreateNewUser(context.Background(), &models.User{
		Email:    "a",
		Name:     "user1",
		Password: "b",
	})
	if err != nil {
		t.Error(err)
	}

	article, err := testRepository.InsertArticle(context.Background(), "", "", "", "", "")
	if err != nil {
		t.Error(err)
	}

	t.Run("Insert", testInsertLabel(testRepository, user, article))
}

func testInsertLabel(testRepository *testRepository, user *models.User, article *models.Article) func(t *testing.T) {
	return func(t *testing.T) {
		label, err := testRepository.InsertLabel(context.Background(), user.ID, article.ID, "Clickbait", "clickbait")
		if err != nil {
			t.Error(err)
		}

		if label.ID == "" {
			t.Errorf("Want label id assigned, got %s", label.ID)
		}
		if got, want := label.UserID, user.ID; got != want {
			t.Errorf("Want label user id %s, got %s", want, got)
		}
		if got, want := label.ArticleID, article.ID; got != want {
			t.Errorf("Want label article id %s, got %s", want, got)
		}
		if got, want := label.Value, "Clickbait"; got != want {
			t.Errorf("Want label value %s, got %s", want, got)
		}
		if got, want := label.Type, "clickbait"; got != want {
			t.Errorf("Want label type %s, got %s", want, got)
		}

		t.Run("Is Labeled", testIsLabeled(testRepository, user.ID, article.ID))
		t.Run("Get", testGetLabel(testRepository, user.ID, article.ID, label.ID))
		t.Run("Get Count", testGetArticleLabelCount(testRepository, article.ID))
		t.Run("Delete", testDeleteLabel(testRepository, label.ID, article.ID))
		t.Run("Get Leaderboard", testGetClickbaitLeaderboard(testRepository, article.ID, user))
		t.Run("Get progress", testGetLabelProgress(testRepository))
	}
}

func testIsLabeled(testRepository *testRepository, userID string, articleID string) func(t *testing.T) {
	return func(t *testing.T) {
		isLabeled, err := testRepository.IsLabeledByUser(context.Background(), articleID, userID, "clickbait")
		if err != nil {
			t.Error(err)
		}

		if !isLabeled {
			t.Errorf("Want is labeled to be true, got %v", isLabeled)
		}
	}
}

func testGetLabel(testRepository *testRepository, userID string, articleID string, labelID string) func(t *testing.T) {
	return func(t *testing.T) {
		labels, err := testRepository.GetLabel(context.Background(), userID, "clickbait")
		if err != nil {
			t.Error(err)
		}

		if got, want := len(labels), 1; got != want {
			t.Errorf("Want get label count %d, got %d", want, got)
		}
		if got, want := labels[0].ID, labelID; got != want {
			t.Errorf("Want label id %s, got %s", want, got)
		}
		if got, want := labels[0].R.Article.ID, articleID; got != want {
			t.Errorf("Want article id %s, got %s", want, got)
		}

		labels, err = testRepository.GetLabel(context.Background(), userID, "summary")
		if err != nil {
			t.Error(err)
		}

		if got, want := len(labels), 0; got != want {
			t.Errorf("Want get label return nothing, got %+v", labels)
		}
	}
}

func testGetArticleLabelCount(testRepository *testRepository, articleID string) func(t *testing.T) {
	return func(t *testing.T) {
		labelCount, err := testRepository.GetArticleLabelCount(context.Background(), articleID, "clickbait")
		if err != nil {
			t.Error(err)
		}

		if got, want := labelCount, int64(1); got != want {
			t.Errorf("Want get label count %d, got %d", want, got)
		}

		labelCount, err = testRepository.GetArticleLabelCount(context.Background(), articleID, "summary")
		if err != nil {
			t.Error(err)
		}

		if got, want := labelCount, int64(0); got != want {
			t.Errorf("Want get label count %d, got %d", want, labelCount)
		}
	}
}

func testDeleteLabel(testRepository *testRepository, labelID string, articleID string) func(t *testing.T) {
	return func(t *testing.T) {
		err := testRepository.DeleteLabel(context.Background(), labelID)
		if err != nil {
			t.Error(err)
		}

		labelCount, err := testRepository.GetArticleLabelCount(context.Background(), articleID, "clickbait")
		if err != nil {
			t.Error(err)
		}

		if got, want := labelCount, int64(0); got != want {
			t.Errorf("Want get label count %d, got %d", want, got)
		}
	}
}

func testGetClickbaitLeaderboard(testRepository *testRepository, articleID string, user1 *models.User) func(t *testing.T) {
	return func(t *testing.T) {
		user2, err := testRepository.CreateNewUser(
			context.Background(),
			&models.User{
				Email:    "b",
				Password: "c",
				Name:     "user2",
			},
		)
		if err != nil {
			t.Error(err)
		}

		insertLabel := func(articleID, userID string) {
			t.Helper()
			if _, err := testRepository.InsertLabel(
				context.Background(),
				userID,
				articleID,
				"",
				"clickbait",
			); err != nil {
				t.Error(err)
			}
		}

		insertLabel(articleID, user1.ID)

		leaderboard, err := testRepository.GetLabelLeaderboard(context.Background(), "clickbait", 10)
		if err != nil {
			t.Error(err)
		}
		log.Println(leaderboard)

		if got, want := len(*leaderboard), 1; got != want {
			t.Errorf("got leaderboard length %d, want %d", got, want)
		}

		if got, want := (*leaderboard)[0].Name, user1.Name; got != want {
			t.Errorf("got user name %s, want %s", got, want)
		}

		if got, want := (*leaderboard)[0].Score, uint32(1); got != want {
			t.Errorf("got leaderboard score %d, want %d", got, want)
		}

		article2, err := testRepository.InsertArticle(context.Background(), "", "", "", "", "")
		if err != nil {
			t.Error(err)
		}

		insertLabel(articleID, user2.ID)
		insertLabel(article2.ID, user2.ID)

		leaderboard, err = testRepository.GetLabelLeaderboard(context.Background(), "clickbait", 10)
		if err != nil {
			t.Error(err)
		}
		log.Println(leaderboard)

		if got, want := len(*leaderboard), 2; got != want {
			t.Errorf("got leaderboard length %d, want %d", got, want)
		}

		if got, want := (*leaderboard)[0].Name, user2.Name; got != want {
			t.Errorf("got user name %s, want %s", got, want)
		}

		if got, want := (*leaderboard)[0].Score, uint32(2); got != want {
			t.Errorf("got leaderboard score %d, want %d", got, want)
		}

		if got, want := (*leaderboard)[1].Name, user1.Name; got != want {
			t.Errorf("got user name %s, want %s", got, want)
		}

		if got, want := (*leaderboard)[1].Score, uint32(1); got != want {
			t.Errorf("got leaderboard score %d, want %d", got, want)
		}

		t.Run("Get label count", testGetLabelCount(testRepository, user1.ID, user2.ID))
	}
}

func testGetLabelProgress(testRepository *testRepository) func(t *testing.T) {
	return func(t *testing.T) {
		labelProgressWeeks, err := testRepository.GetLabelProgress(context.Background(), "", time.Now(), 5)
		if err != nil {
			t.Error(err)
		}

		for _, progress := range labelProgressWeeks[0] {
			if progress.Name == "user1" {
				if got, want := progress.Progress, int64(1); got != want {
					t.Errorf("got progress for user1 %d, want %d", got, want)
				}
			}
			if progress.Name == "user2" {
				if got, want := progress.Progress, int64(2); got != want {
					t.Errorf("got progress for user2 %d, want %d", got, want)
				}
			}
		}
	}
}

func testGetLabelCount(testRepository *testRepository, user1ID, user2ID string) func(t *testing.T) {
	return func(t *testing.T) {
		count1, err := testRepository.GetLabelCount(context.Background(), user1ID)
		if err != nil {
			t.Error(err)
		}
		if got, want := count1, int64(1); got != want {
			t.Errorf("got label count for user1 %d, want %d", got, want)
		}

		count2, err := testRepository.GetLabelCount(context.Background(), user2ID)
		if err != nil {
			t.Error(err)
		}
		if got, want := count2, int64(2); got != want {
			t.Errorf("got label count for user2 %d, want %d", got, want)
		}
	}
}
