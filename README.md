### Adapter Pattern

- go의 디자인 패턴 중 하나
- 서로 다른 구조체 간의 연결을 위한 패턴으로써, interface를 활용한다.
- http 패키지 중 HandlerFunc가 어댑터 패턴을 가지고 있음.

```
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

HandlerFunc *타입*은 ServeHTTP 메소드를 구현하고 있음.

```
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Handler *타입*을 구현하려면 ServeHTTP를 구현해야 한다.

즉, HandlerFunc는 Handler 타입을 구현할 수 있게 도와주는 어댑터인 것.

```
func Main(rw http.ResposeWriter, r *http.Request) {
    rw.Write([]byte("This is Main page"))
}

http.Handle("/", HandleFunc(Main))
```
