package main

import "fmt"

// Старый интерфейс
type OldPrinter interface {
	Print(s string) string
}

// Конкретная реализация старого принтера
type LegacyPrinter struct{}

func (l *LegacyPrinter) Print(s string) string {
	return "Legacy printer: " + s
}

// Новый интерфейс, который ожидает клиент
type NewPrinter interface {
	PrintFormatted(msg string)
}

// Адаптер c использованием OldPrinter
type PrinterAdapter struct {
	OldPrinter OldPrinter
}

// Реализация нового интерфейса с использованием старого
func (p *PrinterAdapter) PrintFormatted(msg string) {
	output := p.OldPrinter.Print(msg)
	fmt.Println(output)
}

func main() {
	adapter := &PrinterAdapter{
		OldPrinter: &LegacyPrinter{},
	}

	adapter.PrintFormatted("Hello from client")
}
