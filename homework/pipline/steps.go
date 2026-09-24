package pipline

import (
	"fmt"
	"sync"
	"time"
)

// Этап передачи в сборку
func Step1(in chan any) chan any {
	out := make(chan any, 3)

	var wg sync.WaitGroup
	go func() {
		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for value := range in {
					deviceOrder, ok := value.(DeviceOrder)
					if !ok {
						fmt.Println("Ошибка при обработке заказа", value)
						continue
					}
					ptrDeviceOrder := &deviceOrder
					fmt.Printf("Заказ %d, сборка\n", deviceOrder.ID)
					device, _ := ptrDeviceOrder.MoveToAssembly()
					out <- device
				}
			}()
		}

		wg.Wait()
		close(out)
	}()

	return out
}

// Этап прошивки
func Step2(in chan any) chan any {
	out := make(chan any, 3)

	// Прошивка для устройства
	firmware := Firmware{
		Version:     "1.1.1",
		PublishDate: time.Now(),
	}

	var wg sync.WaitGroup
	go func() {
		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for value := range in {
					device, ok := value.(Device)
					if !ok {
						fmt.Println("Ошибка прошивки", value)
						continue
					}
					fmt.Printf("Заказ %d, прошивка\n", device.DOID)
					firmware.Update(&device)
					out <- device
				}
			}()
		}

		wg.Wait()
		close(out)
	}()

	return out
}

// Этап тестирования
func Step3(in chan any) chan any {
	out := make(chan any, 3)

	var wg sync.WaitGroup
	go func() {
		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for value := range in {
					device, ok := value.(Device)
					if !ok {
						fmt.Println("Ошибка тестирования", value)
						continue
					}
					ptrDevice := &device
					fmt.Printf("Заказ %d, тестирование\n", device.DOID)
					ptrDevice.Test()
					out <- device
				}
			}()
		}

		wg.Wait()
		close(out)
	}()

	return out
}

// Этап упаковки
func Step4(in chan any) chan any {
	out := make(chan any, 3)

	var wg sync.WaitGroup
	go func() {
		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for value := range in {
					device, ok := value.(Device)
					if !ok {
						fmt.Println("Ошибка упаковки", value)
						continue
					}
					ptrDevice := &device
					fmt.Printf("Заказ %d, упаковка\n", device.DOID)
					resPackage := ptrDevice.Pack()
					out <- resPackage
				}
			}()
		}

		wg.Wait()
		close(out)
	}()

	return out
}
