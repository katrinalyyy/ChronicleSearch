package repository

import (
	"fmt"
	"strings"
)

type ChronicleModel struct {
}

func NewChronicleModel() (*ChronicleModel, error) {
	return &ChronicleModel{}, nil
}

type ChronicleResource struct {
	ID                   int
	Image                string // ключ изображения в MinIO
	ImageKey             string // латинский ключ для MinIO
	Title                string // ОБЯЗАТЕЛЬНО должны быть написаны с заглавной буквы (то есть публичными)
	Author               string
	DateOfCreation       string
	TimeOfAction         string
	Location             string
	DetailedDescription  string
	DetailedSignificance string
	DetailedEditions     string
}

var chronicleResources = []ChronicleResource{ // массив элементов из наших структур
	{
		ID: 1,
		// Image:          "/static/img/povest_vremen_let.png",
		Image:                "http://127.0.0.1:9000/books/povest_vremen_let.png",
		ImageKey:             "povest-vremennykh-let.png",
		Title:                "Повесть временных лет",
		Author:               "Нестор Летописец",
		DateOfCreation:       "1110 - 1118",
		TimeOfAction:         "852 - 1117",
		Location:             "Киевская Русь",
		DetailedDescription:  "«Повесть временных лет» — наиболее ранний из дошедших до нас древнерусских летописных сводов начала XII века. Летопись вобрала в себя в большом количестве материалы сказаний, повестей, легенд, устные поэтические предания о различных исторических лицах и событиях. Охватывает период с библейских времён до 1117 года и считается одним из важнейших источников по истории Древней Руси. Традиционно авторство приписывается монаху Киево-Печерского монастыря Нестору, однако современные исследователи полагают, что летопись создавалась несколькими авторами.",
		DetailedSignificance: "Летопись содержит древнейшие сведения о происхождении Руси, призвании варягов, крещении князя Владимира, походах против Византии и хазар. В ней впервые упоминаются многие города: Новгород, Полоцк, Ростов, Муром, Белоозеро и другие.",
		DetailedEditions:     "Лаврентьевская (1377) - древнейший полный список\n Ипатьевская (начало XV в.) - содержит уникальные сведения\n Радзивилловская (XV в.) - иллюстрированный список",
	},
	{
		ID: 2,
		// Image:          "/static/img/novgorod.png",
		Image:                "http://127.0.0.1:9000/books/troick.png",
		ImageKey:             "novgorod.png",
		Title:                "Новгородская первая летопись",
		Author:               "большинство неизвестно",
		DateOfCreation:       "XIII - XV",
		TimeOfAction:         "1016 - 1471",
		Location:             "Новгород",
		DetailedDescription:  "Новгородская первая летопись — древнерусский летописный свод, отражающий историю Новгорода Великого и северо-западных земель Руси. Создавалась в XIII-XIV веках в Новгороде, представляет собой уникальный источник по истории Новгородской республики. Отличается от киевского летописания более демократическим характером и вниманием к городской жизни, торговле и ремеслам.",
		DetailedSignificance: "Летопись содержит уникальные сведения о политическом устройстве Новгородской республики, взаимоотношениях с князьями, торговых связях с Ганзой и Византией. Важнейший источник по истории борьбы Новгорода с крестоносцами, включая Ледовое побоище 1242 года. Освещает события, связанные с деятельностью Александра Невского, а также монгольское нашествие с новгородской точки зрения.",
		DetailedEditions:     "Синодальный список (XIII-XIV вв.) - древнейшая редакция, частично утрачена\n Комиссионный список (XV в.) - наиболее полный и исправный текст\n Академический список (XV в.) - содержит дополнительные известия",
	},
	{
		ID: 3,
		// Image:          "/static/img/lavrentiv.png",
		Image:                "http://127.0.0.1:9000/books/lavrentiv.png",
		ImageKey:             "lavrentiv.png",
		Title:                "Лаврентьевская летопись",
		Author:               "Лаврентий",
		DateOfCreation:       "1377",
		TimeOfAction:         "1111 - 1305",
		Location:             "Суздаль",
		DetailedDescription:  "Лаврентьевская летопись — древнерусский летописный свод, переписанный в 1377 году монахом Лаврентием по заказу суздальско-нижегородского князя Дмитрия Константиновича. Представляет собой северо-восточную (суздальскую) версию общерусского летописания. Охватывает период с древнейших времен до 1305 года, является одним из важнейших источников по истории Владимиро-Суздальского княжества.",
		DetailedSignificance: "Содержит подробные сведения о становлении Московского княжества, деятельности Андрея Боголюбского, Всеволода Большое Гнездо. Важнейший источник по истории монгольского нашествия, взаимоотношений русских князей с Ордой. Включает уникальную информацию о строительстве городов, храмов, политических событиях северо-восточной Руси.",
		DetailedEditions:     "Лаврентьевский список (1377) - основной список, хранится в РНБ\n Радзивилловская летопись (XV в.) - иллюстрированная версия с миниатюрами\n Московско-Академическая летопись (XV в.) - московская переработка",
	},
	{
		ID: 4,
		// Image:          "/static/img/galic.png",
		Image:                "http://127.0.0.1:9000/books/galic.png",
		ImageKey:             "lavgalicrentiv.png",
		Title:                "Галицко-Волынская летопись",
		Author:               "неизвестны",
		DateOfCreation:       "XIII",
		TimeOfAction:         "1201—1291",
		Location:             "Галич-Волынь",
		DetailedDescription:  "Галицко-Волынская летопись — памятник древнерусского летописания XIII века, освещающий историю юго-западной Руси. Создавалась при дворе галицко-волынских князей, прежде всего Даниила Романовича Галицкого. Отличается высоким литературным уровнем, яркими характеристиками исторических деятелей и подробным описанием военных действий.",
		DetailedSignificance: "Единственный подробный источник по истории Галицко-Волынского княжества в период монгольского нашествия. Содержит уникальные сведения о европейской политике русских князей, отношениях с Польшей, Венгрией, Литвой. Важен для понимания процессов формирования украинской государственности и культуры. Описывает коронацию Даниила Галицкого как короля Руси.",
		DetailedEditions:     "Ипатьевский список (начало XV в.) - основная редакция в составе Ипатьевской летописи\n Хлебниковский список (XVI в.) - дополнительная редакция с вариантами\n Погодинский список (XVII в.) - поздний список с искажениями",
	},
	{
		ID: 5,
		// Image:          "/static/img/pscov.png",
		Image:                "http://127.0.0.1:9000/books/pscov.png",
		ImageKey:             "pscov.png",
		Title:                "Псковские летописи",
		Author:               "неизвестны",
		DateOfCreation:       "XIV - XVI",
		TimeOfAction:         "XIII - XVI",
		Location:             "Псков",
		DetailedDescription:  "Псковские летописи — группа древнерусских летописных памятников XIV-XVI веков, созданных в Пскове. Отражают историю Псковской республики, ее политическое устройство, взаимоотношения с Новгородом, Москвой, Литвой и Ливонским орденом. Характеризуются лаконичностью изложения и точностью в передаче местных событий.",
		DetailedSignificance: "Важнейший источник по истории Псковской вечевой республики, ее демократических институтов. Содержит подробные сведения о борьбе с немецкими рыцарями, осадах Пскова, торговых отношениях. Освещает процесс присоединения Пскова к Московскому государству в 1510 году. Уникальный материал по истории северо-западных рубежей Руси.",
		DetailedEditions:     "Первая Псковская летопись (XV в.) - древнейшая редакция\n Вторая Псковская летопись (XV-XVI вв.) - наиболее подробная\n Третья Псковская летопись (XVI в.) - краткая редакция",
	},
	{
		ID: 6,
		// Image:          "/static/img/troick.png",
		Image:                "http://127.0.0.1:9000/books/troick.png",
		ImageKey:             "troick.png",
		Title:                "Троицкая летопись",
		Author:               "Епифаний Премудрый",
		DateOfCreation:       "XV",
		TimeOfAction:         "XIV - XV",
		Location:             "Москва",
		DetailedDescription:  "Троицкая летопись — московский летописный свод начала XV века, составленный в Троице-Сергиевом монастыре. Погибла в московском пожаре 1812 года, известна по выпискам Н.М. Карамзина. Представляла собой официальную московскую версию русской истории, отражавшую политические интересы московских князей и их стремление к объединению русских земель.",
		DetailedSignificance: "Первый целостный московский летописный свод, обосновывавший права московских князей на общерусское наследство. Содержала подробные сведения о Куликовской битве, деятельности Дмитрия Донского, Сергия Радонежского. Важна для понимания идеологии московского центра, процессов централизации и формирования российской государственности.",
		DetailedEditions:     "Троицкая летопись (начало XV в.) - основной текст (утрачен)\n Симеоновская летопись (XV в.) - близкий к Троицкой текст\n Рогожский летописец (XV в.) - параллельная редакция",
	},
}

func (r *ChronicleModel) GetChronicleResources() ([]ChronicleResource, error) {
	if len(chronicleResources) == 0 {
		return nil, fmt.Errorf("массив пустой")
	}
	return chronicleResources, nil
}

func (r *ChronicleModel) GetChronicleResource(id int) (ChronicleResource, error) {
	// тут у вас будет логика получения нужной услуги, тоже наверное через цикл в первой лабе, и через запрос к БД начиная со второй
	chronicleResources, err := r.GetChronicleResources()
	if err != nil {
		return ChronicleResource{}, err // тут у нас уже есть кастомная ошибка из нашего метода, поэтому мы можем просто вернуть ее
	}

	for _, resource := range chronicleResources {
		if resource.ID == id {
			return resource, nil // если нашли, то просто возвращаем найденный заказ (услугу) без ошибок
		}
	}
	return ChronicleResource{}, fmt.Errorf("услуга не найдена") // тут нужна кастомная ошибка, чтобы понимать на каком этапе возникла ошибка и что произошло
}

func (r *ChronicleModel) GetChronicleResourcesByTitle(title string) ([]ChronicleResource, error) {
	chronicleResources, err := r.GetChronicleResources()
	if err != nil {
		return []ChronicleResource{}, err
	}

	var result []ChronicleResource
	for _, resource := range chronicleResources {
		if strings.Contains(strings.ToLower(resource.Title), strings.ToLower(title)) {
			result = append(result, resource)
		}
	}

	return result, nil
}

type RequestChronicleResearch struct {
	ID          int
	Name        string
	SearchEvent string
}

type ChronicleResearch struct {
	IDRequestResearch int
	IDResource        int
	Quote             string
	IsMatched         bool
}

var requestChronicleResearch = map[RequestChronicleResearch][]ChronicleResearch{
	{ID: 1, Name: "Крещение Руси", SearchEvent: "крести всю землю"}: {
		{
			IDRequestResearch: 1,
			IDResource:        chronicleResources[0].ID,
			Quote:             "Владимир крести всю землю Русскую",
			IsMatched:         true,
		},
		{
			IDRequestResearch: 2,
			IDResource:        chronicleResources[5].ID,
			Quote:             "Епифаний Премудрый написа житие святых",
			IsMatched:         false,
		},
		{
			IDRequestResearch: 3,
			IDResource:        chronicleResources[3].ID,
			Quote:             "Данило князь венча ся короно",
			IsMatched:         false,
		},
	},
}

func (r *ChronicleModel) GetChronicleResearchForRequest(requestID int) (RequestChronicleResearch, []ChronicleResearch, error) {
	for req, research := range requestChronicleResearch {
		if req.ID == requestID {
			return req, research, nil
		}
	}
	return RequestChronicleResearch{}, nil, fmt.Errorf("такой заявки нет")
}
