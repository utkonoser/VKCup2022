### Этот README файл является заготовкой для презентации, само задание можно найти в конце.


### Логика выполнения
1. Запускаю сканирование страниц, создавая очередь из ссылок для корректного обхода в ширину (в виде очереди я использую канал `ChanLinks`). При встрече URL-ссылок изображений, отправляю их на дальнейшую обработку по каналу `ChanPics`.
```go
func processPage(page Page, host string) {
	// какой-то код
	page.Links, page.ImagesUrl = getLinksAndImages(string(body), host)
	// снова какой-то код
	
	for _, imageUrl := range page.ImagesUrl {
		ChanPics <- imageUrl
	}
	for _, link := range page.Links {
		ChanLinks <- link
	}
}

func ProcessingAllPages(host string) {
    // какой-то код
    for link := range ChanLinks {
        page := Page{URL: host + link}
        processPage(page, host)
    }
}
```
2. Параллельно запускается обработка URL-ссылок изображений в другой части программы, которые в свою очередь сначала загружаются, а потом обрабатываются (удаление -Red и -Green составляющих пикселя). Далее уже новое изображение отправляется в соответствующий канал для создания коллажа.
```go
func Collage(imageUrls chan string) image.Image {
    // какой-то код
	go func() {
		defer wg.Done()
		for imageUrl := range imageUrls {
			newImg, err := downloadImage(imageUrl)
			if err != nil {
				log.Fatal("incorrect download images")
			}
			chanImages <- newImg
		}
	}()
	// еще немного какого-то кода
}
func downloadImage(url string) (image.Image, error) {
	// какой
    resp, err := http.Get(url)
	// то
    img, err := png.Decode(resp.Body)
	// код 
	result := removeRedGreen(img)
	// тут
    return result, nil
}
```
3. Считываю канал с результирующими изображениями и добавляю их в коллаж. В конце создаю конечное PNG-изображение.
```go
func Collage(imageUrls chan string) image.Image {
	// какой-то код
	cols := 32
	rows := len(chanImages) / cols
	if len(chanImages)%32 != 0 {
		rows++
	}

	collage := image.NewRGBA(image.Rect(0, 0, cols*512, rows*512))

	i := 0
	for src := range chanImages {
		x := (i % cols) * 512
		y := (i / cols) * 512
		i++
		draw.Draw(collage, image.Rect(x, y, x+512, y+512), src, image.Point{}, draw.Src)
	}
	// какой-то код
	return collage
}
```

### Пример работы:

```shell
ni@ni-asus:~/GolandProjects/Final/cmd/solution$ go run .  http://localhost:8080 result.png
2023/02/07 13:02:14 Application start...
2023/02/07 13:02:14 Loading images and changing pixels in them...
2023/02/07 13:03:18 End of changing images in: 1m4.326287971s
2023/02/07 13:03:18 Creating a collage...
2023/02/07 13:03:19 The creation of the collage took: 598.424684ms
2023/02/07 13:03:19 PNG image encoding...
2023/02/07 13:03:30 Application running time 1m16.214408021s 
```


### Какие решения я использовал
* Для кеширования посещенных страниц во время обхода графа сайта, я использовал `sync.Map`. Также во второй раз `sync.Map` использую для кеширования уже измененных изображений, например:
```go
func downloadImage(url string) (image.Image, error) {
	if val, ok := imagesCache.Load(url); !ok {
		// делаю запрос по URL
		result := removeRedGreen(img)
		imagesCache.Store(url, result)
		return result, nil
	} else {
		return val.(image.Image), nil
	}
}
```
*  С помощью каналов добился параллельного исполнения сканирования страниц и обработки изображений.

### Задание

Вам дан сайт. Чтобы открыть его на порту 8080, скачайте докер образ vkcup2022/golang:latest и запустите его строчкой

docker run -it --rm -d -p 8080:80 --name web vkcup2022/golang:latest

На каждой его странице может быть одно или несколько изображений размером 512×512 пикселей в формате PNG, а также одна или несколько ссылок. Обойдите граф сайта в ширину.

В каждой из предложенных картинок замените Red- и Green-составляющие каждого пикселя на 0 таким образом, чтобы осталась только Blue-составляющая.

После этого объедините все фотографии в единый коллаж именно в том порядке, в котором они встретились вам при обходе в ширину. Ширина коллажа фиксированная: 32 изображения (512×32 = 16384 пикселя). Заполните коллаж слева направо, а потом сверху вниз.

Обратите внимание: на итоговом сайте будет очень большое количество изображений. Но ваше решение должно использовать не более 2 Гбайт оперативной памяти.

Ваша задача — написать как можно более эффективное решение. Чем меньше времени потребуется вашему решению на запуск, тем выше будет ваше место на соревновании. Мы будем запускать ваше решение строкой

./solution url result.png

Вам нужно будет загрузить задание и презентацию до 10:00 по Москве 4 февраля.

Кроме исходного кода выполненного задания, вам нужно будет подготовить и загрузить четырёхминутную презентацию, рассказывающую о вашем решении.

    Включите в презентацию три примера результата запуска вашей программы.
    Расскажите, какие интересные технические решения вы использовали? Не рассказывайте про очевидное.
    Что бы вы доделали, если бы нужно было в ближайшее время запускать этот продукт на пользователей?

Презентация и защита проектов начнётся 5 февраля в 11:00. Позже мы опубликуем информацию о порядке выступлений и свяжемся с вами, чтобы запланировать технический прогон.
