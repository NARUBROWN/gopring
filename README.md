<p align="center">
  <img src="assets/gospring-logo.png" alt="GoSpring 로고" width="420" />
</p>

# GoSpring

GoSpring은 Go + Spring 스타일을 간단히 흉내 낸 학습용 미니 프레임워크입니다. Echo를 서블릿 컨테이너처럼 사용하고, 자체 DI 컨테이너와 디스패처를 통해 컨트롤러 매핑/인자 바인딩/반환값 처리 흐름을 구현합니다.

## 무엇을 하는 프로젝트인가요?

- `ApplicationContext` 기반의 싱글턴 DI 컨테이너와 순환 의존성 체크
- 컨트롤러의 `Mappings()` 메서드로 라우팅 메타데이터 수집
- 요청 경로 매칭 및 Path Variable/Query Param 바인딩
- 반환 타입에 따라 응답 처리 (string, struct(JSON), error)

## 동작 흐름

1. `app.Bootstrap()`에서 `ApplicationContext`를 생성하고 빈을 등록합니다.
2. `web.Dispatcher`가 컨트롤러 매핑을 스캔해 핸들러 목록을 구성합니다.
3. Echo가 모든 요청을 디스패처로 위임합니다.
4. 디스패처가 경로/메서드를 매칭하고, 인자를 리졸빙해 핸들러 메서드를 실행합니다.
5. 반환값 타입에 따라 응답을 만들어 반환합니다.

## 실행 방법

```bash
go run .
```

```bash
# Query Param 바인딩 예시
curl "http://localhost:8080/users?name=kim"

# Path Variable 바인딩 예시
curl "http://localhost:8080/users/10"
```

## 패키지 구조

- `app/`: 부트스트랩, 빈 등록
- `context/`: 간단한 DI 컨테이너 (`ApplicationContext`)
- `web/`: 디스패처, 매핑/리졸버/리턴 핸들러
- `example/`: 예제 컨트롤러/서비스/리포지토리

## 확장 포인트

- `ArgumentResolver`를 추가해 바인딩 규칙 확장
- `ReturnValueHandler`를 추가해 반환 타입 확장
- `Mappings()` 규약을 확장해 어노테이션 스타일 메타데이터로 발전 가능

## 참고

이 프로젝트는 Spring MVC 흐름(DispatcherServlet, HandlerMethod, ArgumentResolver, ReturnValueHandler)을 Go/Echo 위에서 단순화한 학습용 샘플입니다.
