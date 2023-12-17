package service

const (
	clientIPAPI = "api.bigdatacloud.net/data/client-ip"
)

type IPService interface {
	GetPublicIP() (string, error)
}

type ipservice struct{}

func NewIPService() IPService {
	return &ipservice{}
}

func (s *ipservice) GetPublicIP() (string, error) {
	// resp, err := http.Get("https://" + clientIPAPI)
	// if err != nil {
	// 	return "", err
	// }

	// if resp.StatusCode != 200 {
	// 	return "", errors.New(fmt.Sprintf("Resp was a non 200 code: %s", resp.Status))
	// }

	// b, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return "", err
	// }

	// // resp interface
	// type apiResp struct{}
	// var d *apiResp

	// if err := json.Unmarshal(b, d); err != nil {
	// 	return "", err
	// }

	return "192.168.4.111", nil
}
