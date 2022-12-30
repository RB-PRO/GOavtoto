package avtotoGo

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Метод SearchGetParts2
// Описание: Предназначен для получения результатов поиска запчастей по коду на сервере AvtoTO. Расширенная версия, выдает статус ответа.

//const SearchStatus = [...]string{"Неверно указан ID процесса ProcessSearchId", "Запрос не найден", "Запрос в обработке", "Ошибка данных", "Результат получен"}

// Структура запроса метода SearchGetParts2
type SearchGetParts2Request struct {
	ProcessSearchId string `json:"ProcessSearchId"` // Уникальный идентификатор процесса поиска (тип: строка).
	Limit           int    `json:"Limit"`           // необязательный параметр, орграничение на количество строк в выдаче (тип: целое).
}

// Структура ответа метода SearchGetParts2
type SearchGetParts2Response struct {
	// Список запчастей, найденных по запросу - индексированный массив с упорядоченными целочисленными ключами, начиная с 0.
	// Каждый элемент этого массива содержит информацию о конкретной детали и представляет из себя ассоциативный массив.
	// Свойства детали:
	Parts []struct {
		Code      string `json:"Code"`      // [*] Код детали
		Manuf     string `json:"Manuf"`     // [*] Производитель
		Name      string `json:"Name"`      // [*] Название
		Price     int    `json:"Price"`     // Цена
		Storage   string `json:"Storage"`   // [*] Склад
		Delivery  string `json:"Delivery"`  // [*] Срок доставки
		MaxCount  string `json:"MaxCount"`  // [*] Максимальное количество для заказа, остаток по складу. Значение "-1" - означает "много" или "неизвестно"
		BaseCount string `json:"BaseCount"` // [*] Кратность заказа

		StorageDate     string `json:"StorageDate"`     // Дата обновления склада // В случае, когда SearchStatus = 4 (Результат получен)
		DeliveryPercent int    `json:"DeliveryPercent"` // Процент успешных закупок из общего числа заказов // В случае, когда SearchStatus = 4 (Результат получен)
		BackPercent     int    `json:"BackPercent"`     // Процент удержания при возврате товара (при отсутствии возврата поставщику возвращается значение "-1") // В случае, когда SearchStatus = 4 (Результат получен)

		AvtotoData struct { // Массив со след. элементами:
			PartId int `json:"PartId"` // [*] Номер запчасти в списке результата поиска
		} `json:"AvtotoData"`
	} `json:"Parts"`

	Info struct {
		Errors       []string `json:"Errors"`       // Массив ошибок, возникший в процессе поиска
		SearchStatus int      `json:"SearchStatus"` // Информация о статусе процесса на сервере AvtoTO. Возможные варианты значений:
		SearchID     string   `json:"SearchId"`     // Уникальный идентификатор запроса поиска, возвращается в случае удачного поиска
	} `json:"Info"`
	// [*] — эти данные необходимо сохранить в Вашей системе, в дальнейшем они понадобятся для добавления запчастей в корзину
}

// Получить данные по методу SearchGetParts2
func (SearchGetParts2Req SearchGetParts2Request) SearchGetParts2() (SearchGetParts2Response, error) {

	// Ответ от сервера
	var responseSearchGetParts2 SearchGetParts2Response

	// Подготовить данные для загрузки
	bytesRepresentation, responseError := json.Marshal(SearchGetParts2Req)
	if responseError != nil {
		return responseSearchGetParts2, responseError
	}

	// Отправить данные
	body, responseError := HttpPost(bytesRepresentation, "SearchGetParts2")
	if responseError != nil {
		return responseSearchGetParts2, responseError
	}

	fmt.Println(string(body))
	fmt.Println()

	// Распарсить данные
	responseError = responseSearchGetParts2.SearchGetParts2_UnmarshalJson(body)

	return responseSearchGetParts2, responseError
}

// Метод для SearchGetParts2, который преобразует приходящий ответ в структуру
func (responseSearchGetParts2 *SearchGetParts2Response) SearchGetParts2_UnmarshalJson(body []byte) error {
	responseError := json.Unmarshal(body, &responseSearchGetParts2)
	if responseError != nil {
		return responseError
	}

	if len(responseSearchGetParts2.Info.Errors) != 0 {
		return errors.New(responseSearchGetParts2.Info.Errors[0])
	}
	return nil
}

/*

func SearchGetParts2Data(data ProcessSearchId) []byte {
	method := "POST"
	dts_json_usr, err_json_usr := json.Marshal(data)
	if err_json_usr != nil {
		fmt.Println(err_json_usr)
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("action", "SearchGetParts2")
	_ = writer.WriteField("data", string(dts_json_usr))
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	return body
}
*/
