<H3>단일 책임원칙</H3>

```
type FinanceReport struct{
    report string
}

func (r *FinanceReport) SendReport(email string){
}

type MarketingReport struct{
    report string
}

func (r *MarketingReport) SendReport(email string){
}
```

- SendReport 라는 동일 기능을 이렇게 구현할 이유가 없다

```
type Report interface{
    Report() string
}

type FinanceReport struct{
    report string
}

func (r *FinanceReport) Report() string{
    return r.report
}

type ReportSender struct{

}

func (r *ReportSender) SendReport(report Report){

}
```

- 이렇게 코드를 작성하면 MarketingReport 도 동일한 코드를 재활용 가능

<H3>개방-폐쇄 원칙</H3>
```
type ReportSender interface{
    Send(r *Report)
}

type EmailSender struct{

}

func (e *EmailSender) Send(r *Report){

}
```

- if 혹은 switch 를 사용하는 것보다 객체 각자 책임지는 방식에서 우월하고, switch 의 코드를 그대로 복사 붙여넣는 반복 코드를 작성할 이유가 없다

<H3>리스코프 치환 원칙</H3>
- 인터페이스를 구현하는 구조체는 해당 인터페이스에 지정된 메소드를 무조건 사용할 수 있어야 한다

<H3>인터페이스 분리 원칙</H3>
- 불필요한 메소드를 기능에 따라 분리하여 해당 인터페이스에서만 책임질 수 있도록 한다

<H3>의존 관계 역전 원칙</H3>
1. 상위 모듈은 하위 모듈에 의존하지 않는다
    
    -> 하위 모듈도 추상화 시켜서 상호간 의존도를 낮추면 된다
    
    -> ex) 키보드를 입력 장치로 추상화 시키면 스피커, 목소리 등등도 입력장치가 된다


2. 추상화 시킨 모듈을 생성하여 구조체간 종속을 분리 시킨다


