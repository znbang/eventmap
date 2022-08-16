package bookservice

import (
	"errors"
	"log"
	"time"

	"github.com/znbang/eventmap/pkg/crawlerservice"
	"github.com/znbang/eventmap/pkg/uuid"
)

func (s *BookService) HandleBookJob() {
	for {
		time.Sleep(5 * time.Second)

		jobs := make([]BookJob, 0)
		if err := s.bookJobRepository.Find100Pending(&jobs); err != nil {
			log.Println(err)
		}

		for _, job := range jobs {
			if err := s.markTaskRunning(job); err != nil {
				log.Println(err)
			}
			if err := s.handleBookJob(job); err != nil {
				log.Println(err)
				if err = s.markTaskDone(job, err.Error()); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func (s *BookService) handleBookJob(job BookJob) error {
	var (
		book    Book
		chapter Chapter
		url     string
		page    int
	)

	// 取得第一章下載 URL
	if err := s.bookRepository.FindById(&book, job.BookID); err == nil {
		url = book.URL
	} else if errors.Is(err, ErrNoSuchBook) {
		return errors.New("找不到書本。")
	} else {
		return err
	}

	// 取得最末章下載 URL 和頁碼
	if err := s.chapterRepository.FindLastChapterByBookId(&chapter, job.BookID); err == nil {
		url = chapter.URL
		page = chapter.Page
	} else if errors.Is(err, ErrNoSuchChapter) {
	} else {
		return err
	}

	crawler := crawlerservice.New()
	skipFirstResult := page > 0

	for url != "" {
		var result crawlerservice.Result
		if err := crawler.Download(&result, url); err != nil {
			return err
		}

		if skipFirstResult {
			// 如果不是從第一章開始下載，忽略第一個下載結果
			skipFirstResult = false
			url = result.MoreURL
		} else if result.Title == "" {
			if err := s.updateBookJobMessage(job.ID, "沒有標題。"); err != nil {
				return err
			}
			time.Sleep(30 * time.Second)
		} else if result.Body == "<empty>" {
			if err := s.updateBookJobMessage(job.ID, "沒有內容。"); err != nil {
				return err
			}
			time.Sleep(30 * time.Second)
		} else {
			page = page + 1

			if err := s.chapterRepository.Create(&Chapter{
				ID:     uuid.ULID(),
				BookID: book.ID,
				URL:    url,
				Page:   page,
				Title:  result.Title,
				Body:   result.Body,
			}); err != nil {
				return err
			}

			if err := s.updateBookJobMessage(job.ID, result.Title); err != nil {
				return err
			}

			book.UpdatedAt = time.Now()
			if err := s.bookRepository.Save(&book); err != nil {
				return err
			}

			url = result.MoreURL
		}

		// 檢查是否被使用者停止
		if err := s.bookJobRepository.FindById(&job, job.ID); err != nil {
			return err
		} else if job.IsDone() {
			return nil
		}

		if err := EventBus.Post(Event{job.ID, result.Title, StatusRunning}); err != nil {
			return err
		}

		time.Sleep(5 * time.Second)
	}

	return s.markTaskDone(job, "")
}

func (s *BookService) updateBookJobMessage(jobId string, message string) error {
	var job BookJob

	if err := s.bookJobRepository.FindById(&job, jobId); err != nil {
		return err
	}

	job.Message = message

	return s.bookJobRepository.Save(&job)
}

func (s *BookService) markTaskRunning(job BookJob) error {
	job.Status = StatusRunning
	job.Message = ""
	if err := EventBus.Post(Event{job.ID, "", StatusRunning}); err != nil {
		return err
	}
	return s.bookJobRepository.Save(&job)
}

func (s *BookService) markTaskDone(job BookJob, message string) error {
	job.Status = StatusDone
	job.Message = message
	if err := EventBus.Post(Event{job.ID, message, StatusDone}); err != nil {
		return err
	}
	return s.bookJobRepository.Save(&job)
}
