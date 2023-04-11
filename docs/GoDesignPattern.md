# Go设计模式
### 1.创建型模式
+ **单例模式**
```go
// 饿汉方式
package singleton

type singleton struct {}

var ins *singleton = &singleton{}

func GetInsOr() *singleton {
	return &singleton
}
```
```go
package singleton

import "sync"

type singleton struct {}

var ins *singleton
var once sync.Once

func GetInsOr() *singleton {
	once.Do(func() {
		ins = &singleton{}
    })
	return ins
}
```
+ 简单工厂模式
+ 抽象工厂模式
+ 工厂方法模式
```go

```
+ 建造者模式
+ 原型模式

### 2.结构型模式
+ 访问者模式
+ **模板模式**
```go
package template

import "fmt"

type Cooker interface {
    fire()
	cooke()
	outfire()
}

type CookMenu struct {}

func (CookMenu) fire() {
	fmt.Println("开火")
}

func (CookMenu) cooke() {
    
}

func (CookMenu) outfire()  {
    fmt.Println("关火")
}

func doCook(cooker Cooker) {
	cooker.fire()
	cooker.cooke()
	cooker.outfire()
}

type Xi struct {
    CookMenu
}

func (*Xi) cooke()  {
    fmt.Println("xihongsi")
}

type JiDan struct {
	CookMenu
}

func (*JiDan) cooke()  {
	fmt.Println("JiDan")
}
```
测试用例
```go
package template
func TestTemplate (t *testing.T) {
	xi := &Xi{}
	doCook(xi)
	
	jiDan := &JiDan{}
	doCook(jiDan)
}

```
+ **策略模式**
```go
package strategy

// IStrategy 策略类
type IStrategy interface {
    do(int,int) int
}
// add 策略实现： 加
type add struct {}

func (*add) do(a,b int) int {
    return a + b
}
// reduce 策略实现：减
type reduce struct {}

func (*reduce) do(a,b int) int {
	return a - b
}
// Operator 具体策略执行者
type Operator struct {
    strategy IStrategy
}
// setStrategy 设置策略
func (receiver *Operator) setStrategy(strategy IStrategy)  {
    receiver.strategy = strategy
}
// calculate 调用策略中的方法
func (receiver *Operator) calculate(a,b int) int {
    return receiver.strategy.do(a,b)
}
```
+ 状态模式
+ 观察者模式
+ 备忘录模式
+ 中介者模式
+ 迭代器模式
+ 解释器模式
+ 命令模式
+ 责任链模式

### 3.行为型模式
+ 适配器模式
+ 桥接模式
+ 组合模式
+ 装饰模式
+ 外观模式
+ 享元模式
+ **代理模式**
```go
package proxy

import "fmt"

type Seller interface {
    sell(name string)
}
// Station 火车站
type Station struct {
    stock int
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
        station.stock --
		fmt.Println("卖出去一张票了")
	}else {
		fmt.Println("票以售空")
    }
}
// StationProxy 火车站代理站
type StationProxy struct {
    station *Station
}

func (proxy *StationProxy) sell(name string)  {
	if proxy.station.stock > 0 {
        proxy.station.stock --
		fmt.Println("Selling")
	}else {
		fmt.Println("票以售空")
    }
}
```
+ **选项模式**
```go
package options

import "time"

type Connetion struct {
    addr string
	cache bool
	timeout time.Duration
}

const (
	defaultTimeout = 10
	defaultCache = false
)

type options struct{
	timeout time.Duration
	cache bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options)  {
    f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
    })
}

func WithCache(cache bool) Option {
	return optionFunc(func(o *options) {
		o.cache = cache
    })
}

func NewConnect(addr string, opts ...Option) (*Connetion, error) {
	options := options{
		timeout: defaultTimeout,
		cache: defaultCache,
    }
	for _, o := range opts {
        o.apply(&options)
	}
	
	return &Connetion{
		addr: addr,
		cache: options.cache,
	    timeout: options.timeout,	
	}
}
```

## SOLID原则
+ 单一功能原则
+ 开闭原则
+ 里氏替换原则
+ 依赖倒置原则
+ 接口分离原则