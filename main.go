package main

import (
	"log"
	"time"

	tb "gopkg.in/telebot.v3"
)

func main() {
	pref := tb.Settings{
		Token:  "8267289581:AAE7M036wD_bPSbkHvGoJ8tsIKFLKa5C2tA",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

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
bot.Handle("/find", func(c tb.Context) error {
	// Тестовые параметры: температура -2, влажность 70, снег "old", трасса "hard", стиль "classic"
	recommendations := FilterWaxes(-2, 70, "old", "hard", "classic")
	
	if len(recommendations) == 0 {
		return c.Send("Нет подходящих мазей для таких условий.")
	}
	
	msg := "Подходящие мази:\n"
	for _, r := range recommendations {
		msg += "— " + r.Name + "\n"
	}
	return c.Send(msg)
})





	log.Println("Бот запущен...")
	bot.Start()
}
