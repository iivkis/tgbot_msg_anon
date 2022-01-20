package replicas

import "fmt"

func Get(title string, format ...interface{}) string {
	t, ok := texts[title]
	if !ok {
		fmt.Printf("[Replicas:error] отсутсвует заголовок `%s`", title)
	}

	return fmt.Sprintf(t, format...)
}
