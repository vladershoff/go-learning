package main

import (
	"fmt"
	"homework/pipline"
	"time"
)

func main() {
	// создаем воркера, который будет генерировать 1 заказ в секунду
	// настраиваем пайплайн со следующими шагами
	// логируем заказ
	// передаем в сборку (400 мс): Заказ -> Девайс
	// передаем на прошивку (200 мс): Девайс -> Девайс
	// передаем на тестирование (100 мс): Девайс -> Девайс
	// передаем на упаковку (100 мс): Девайс -> Пакет
	// радуемся жизни, пока

	// Создаем pln и настраиваем этапы
	pln := pipline.NewPipline()
	pln.AddStep(pipline.Step1)
	pln.AddStep(pipline.Step2)
	pln.AddStep(pipline.Step3)
	pln.AddStep(pipline.Step4)

	// Запускаем pipline
	out := pln.Run()

	go func() {
		for i := 1; i <= 100; i++ {
			time.Sleep(time.Second)
			deviceOrder := pipline.DeviceOrder{
				ID:        i,
				CreatedAt: time.Now(),
				Model:     "Apple iPhone 18 Pro Max 256GB, Burgundy",
			}
			pln.PushOrder(deviceOrder)
		}
		pln.Close()
	}()

	for i := range out {
		resPackage := i.(pipline.Package)
		fmt.Printf("Заказ %d, завершена сборка\n", resPackage.Device.DOID)
	}
}
