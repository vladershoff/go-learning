package pipline

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
)

// Заказ на устройство
type DeviceOrder struct {
	ID        int
	CreatedAt time.Time
	Model     string
}

// Собирает девайс из заказа
func (d *DeviceOrder) MoveToAssembly() (Device, error) {
	// Задержка на сборку
	// time.Sleep(400 * time.Millisecond)
	time.Sleep(time.Duration(rand.IntN(1001)) * time.Millisecond)

	// Генерируем идентификатор UUID v7
	id, err := uuid.NewV7()
	if err != nil {
		return Device{}, fmt.Errorf("не удалось сгенерировать id в MoveToAssembly: %w", err)
	}

	// Возвращаем девайс из заказа
	device := Device{
		ID:          id.String(),
		Model:       d.Model,
		DOCreatedAt: d.CreatedAt,
		DOID:        d.ID,
	}
	return device, nil
}

// Прошивка
type Firmware struct {
	Version     string
	PublishDate time.Time
}

// Выполняет прошивку девайса
func (f *Firmware) Update(device *Device) {
	// Задержка на прошивку
	// time.Sleep(200 * time.Millisecond)

	time.Sleep(time.Duration(rand.IntN(1001)) * time.Millisecond)

	// Сохраняем в девайсе данные прошивки
	device.FW = *f
	device.FWUploadedAt = time.Now()
}

// Устройство
type Device struct {
	ID           string // UUID v7
	Model        string
	DOCreatedAt  time.Time
	DOID         int
	AssembledAt  time.Time
	FW           Firmware
	FWUploadedAt time.Time
	QCPassed     bool
	QCAt         time.Time
}

// Тестирует девайс
func (d *Device) Test() {
	// Задержка на тестирование
	// time.Sleep(100 * time.Millisecond)
	time.Sleep(time.Duration(rand.IntN(1001)) * time.Millisecond)

	// Сохраняем данные о тестировании
	d.QCPassed = true
	d.QCAt = time.Now()
}

// Упаковать девайс
func (d *Device) Pack() Package {
	// Задержка на упаковку
	// time.Sleep(100 * time.Millisecond)
	time.Sleep(time.Duration(rand.IntN(1001)) * time.Millisecond)

	// Сохраняем данные об упаковке
	d.AssembledAt = time.Now()

	// Формируем отчет об упаковке
	report := Report{
		DeviceID:      d.ID,
		DeviceModel:   d.Model,
		AssembledAt:   d.AssembledAt,
		AssembledBy:   "",
		FWVersion:     d.FW.Version,
		FWPublishDate: d.FW.PublishDate,
		FWUploadedAt:  d.FWUploadedAt,
		QCPassed:      d.QCPassed,
		QCAt:          d.QCAt,
		TotalTime:     time.Duration(d.AssembledAt.Sub(d.DOCreatedAt)),
	}

	return Package{
		Device: *d,
		Report: report,
	}
}

// Отчет
type Report struct {
	DeviceID      string
	DeviceModel   string
	AssembledAt   time.Time
	AssembledBy   string
	FWVersion     string
	FWPublishDate time.Time
	FWUploadedAt  time.Time
	QCPassed      bool
	QCAt          time.Time
	TotalTime     time.Duration
}

// Пакет, Упаковка
type Package struct {
	Device Device
	Report Report
}
