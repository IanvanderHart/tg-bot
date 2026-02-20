package main

import tb "gopkg.in/telebot.v3"

// Клавиатура для выбора стиля катания
func styleKeyboard() *tb.ReplyMarkup {
	menu := &tb.ReplyMarkup{ResizeKeyboard: true}
	btnClassic := menu.Text("⛷ Классика")
	btnSkate := menu.Text("🏂 Конёк")
	btnAny := menu.Text("⚡ Любой")
	menu.Reply(
		menu.Row(btnClassic, btnSkate),
		menu.Row(btnAny),
	)
	return menu
}

// Клавиатура для выбора типа снега
func snowKeyboard() *tb.ReplyMarkup {
	menu := &tb.ReplyMarkup{ResizeKeyboard: true}
	btnFresh := menu.Text("❄️ Свежий")
	btnOld := menu.Text("🗻 Старый")
	btnWet := menu.Text("💧 Мокрый")
	btnAny := menu.Text("⚡ Любой")
	menu.Reply(
		menu.Row(btnFresh, btnOld),
		menu.Row(btnWet, btnAny),
	)
	return menu
}

// Клавиатура для выбора состояния трассы
func trackKeyboard() *tb.ReplyMarkup {
	menu := &tb.ReplyMarkup{ResizeKeyboard: true}
	btnHard := menu.Text("🧊 Жёсткая")
	btnSoft := menu.Text("🌨 Мягкая")
	btnIcy := menu.Text("🪞 Лёд")
	btnAny := menu.Text("⚡ Любая")
	menu.Reply(
		menu.Row(btnHard, btnSoft),
		menu.Row(btnIcy, btnAny),
	)
	return menu
}

// Клавиатура для подтверждения выбора
func confirmKeyboard() *tb.ReplyMarkup {
	menu := &tb.ReplyMarkup{ResizeKeyboard: true}
	btnYes := menu.Text("✅ Да")
	btnNo := menu.Text("❌ Заново")
	menu.Reply(menu.Row(btnYes, btnNo))
	return menu
}
