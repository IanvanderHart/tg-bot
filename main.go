package main

import (
	"log"
	//"os"
        "time"
        "fmt"

	tb "gopkg.in/telebot.v3"
)

type UserState struct {
	Step     string // "style", "snow", "track", "confirm"
	Style    string
	Snow     string
	Track    string
}

var userStates = make(map[int64]*UserState) // хранилище состояний

func main() {
	pref := tb.Settings{
		Token:  "8267289581:AAE7M036wD_bPSbkHvGoJ8tsIKFLKa5C2tA",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}


//insertion 20260220 14:26
// Начало подбора
bot.Handle("/find", func(c tb.Context) error {
	userID := c.Sender().ID
	userStates[userID] = &UserState{Step: "style"}
	return c.Send("Выбери стиль катания:", styleKeyboard())
})

// Обработка текстовых сообщений (кнопок)
bot.Handle(tb.OnText, func(c tb.Context) error {
	userID := c.Sender().ID
	state, exists := userStates[userID]
	if !exists {
		return nil // пользователь не в режиме подбора
	}

	text := c.Text()

	switch state.Step {
	case "style":
		switch text {
		case "⛷ Классика":
			state.Style = "classic"
		case "🏂 Конёк":
			state.Style = "skate"
		case "⚡ Любой":
			state.Style = "any"
		default:
			return c.Send("Пожалуйста, выбери стиль кнопками.")
		}
		state.Step = "snow"
		return c.Send("Теперь выбери тип снега:", snowKeyboard())

	case "snow":
		switch text {
		case "❄️ Свежий":
			state.Snow = "fresh"
		case "🗻 Старый":
			state.Snow = "old"
		case "💧 Мокрый":
			state.Snow = "wet"
		case "⚡ Любой":
			state.Snow = "any"
		default:
			return c.Send("Пожалуйста, выбери снег кнопками.")
		}
		state.Step = "track"
		return c.Send("Выбери состояние трассы:", trackKeyboard())

	case "track":
		switch text {
		case "🧊 Жёсткая":
			state.Track = "hard"
		case "🌨 Мягкая":
			state.Track = "soft"
		case "🪞 Лёд":
			state.Track = "icy"
		case "⚡ Любая":
			state.Track = "any"
		default:
			return c.Send("Пожалуйста, выбери трассу кнопками.")
		}
		state.Step = "confirm"
		msg := fmt.Sprintf("Твой выбор:\nСтиль: %s\nСнег: %s\nТрасса: %s\n\nПодобрать мазь?", state.Style, state.Snow, state.Track)
		return c.Send(msg, confirmKeyboard())

	case "confirm":
		switch text {
		case "✅ Да":
			// Здесь будем вызывать FilterWaxes с параметрами из state
			// Пока просто заглушка
			delete(userStates, userID)
			return c.Send("Ищу подходящие мази... (скоро заработает)")
		case "❌ Заново":
			delete(userStates, userID)
			return c.Send("Подбор отменён. Начни заново с /find")
		default:
			return c.Send("Пожалуйста, подтверди или отмени.")
		}
	}
	return nil
})

//end insertion 20260220 14:26



	// Команда /start
	bot.Handle("/start", func(c tb.Context) error {
		return c.Send("Привет! Я лыжный бот-помощник. Напиши /help, чтобы узнать, что я умею.")
	})

	// Команда /help
	bot.Handle("/help", func(c tb.Context) error {
		helpText := `Я помогу подобрать лыжную мазь.
Используй команду /recommend, чтобы начать подбор.
Пока я только учусь, но скоро буду спрашивать:
- Температуру
- Влажность
- Качество снега
- Состояние трассы
А потом рекомендовать мазь для классики или конька.`
		return c.Send(helpText)
	})

	// Команда /recommend (заглушка)
	bot.Handle("/recommend", func(c tb.Context) error {
		return c.Send("Пока я только учусь подбирать мази. Скоро здесь появится анкета с вопросами.")
	})

	// Обработчик любого текста (если не команда)
	bot.Handle(tb.OnText, func(c tb.Context) error {
		// Просто игнорируем, чтобы бот не отвечал на каждый чих
		return nil
	})






// Временная команда для тестирования подбора
//bot.Handle("/find", func(c tb.Context) error {
	// Тестовые параметры: температура -2, влажность 70, снег "old", трасса "hard", стиль "classic"
////	recommendations := FilterWaxes(-2, 70, "old", "hard", "classic")
//	
//	if len(recommendations) == 0 {
//		return c.Send("Нет подходящих мазей для таких условий.")
//	}
//	
//	msg := "Подходящие мази:\n"
//	for _, r := range recommendations {
//		msg += "— " + r.Name + "\n"
//	}
//	return c.Send(msg)
//})





	log.Println("Бот запущен...")
	bot.Start()
}

