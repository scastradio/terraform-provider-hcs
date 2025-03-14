package snapshots

import (
	"fmt"
	"time"

	"github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/sdk/huaweicloud"
)

type JobResponse struct {
	JobID string `json:"job_id"`
}

type JobStatus struct {
	Status     string                 `json:"status"`
	Entities   map[string]interface{} `json:"entities"`
	JobID      string                 `json:"job_id"`
	JobType    string                 `json:"job_type"`
	BeginTime  string                 `json:"begin_time"`
	EndTime    string                 `json:"end_time"`
	ErrorCode  string                 `json:"error_code"`
	FailReason string                 `json:"fail_reason"`
}

type JobResult struct {
	golangsdk.Result
}

func (r JobResult) ExtractJobResponse() (*JobResponse, error) {
	job := new(JobResponse)
	err := r.ExtractInto(job)
	return job, err
}

func (r JobResult) ExtractJobStatus() (*JobStatus, error) {
	job := new(JobStatus)
	err := r.ExtractInto(job)
	return job, err
}

func WaitForJobSuccess(client *golangsdk.ServiceClient, secs int, jobID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		job := new(JobStatus)
		_, err := client.Get(jobURL(client, jobID), &job, nil)
		time.Sleep(5 * time.Second)
		if err != nil {
			return false, err
		}

		if job.Status == "SUCCESS" {
			return true, nil
		}
		if job.Status == "FAIL" {
			err = fmt.Errorf("job failed with code %s: %s", job.ErrorCode, job.FailReason)
			return false, err
		}

		return false, nil
	})
}

func GetJobEntity(client *golangsdk.ServiceClient, jobID string, label string) (interface{}, error) {
	job := new(JobStatus)
	_, err := client.Get(jobURL(client, jobID), &job, nil)
	if err != nil {
		return nil, err
	}

	if job.Status == "SUCCESS" {
		if e, ok := job.Entities[label]; ok {
			return e, nil
		}
	}

	return nil, fmt.Errorf("unexpected conversion error in GetJobEntity")
}
