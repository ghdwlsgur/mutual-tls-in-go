### 1. mTLS (Mutual Transport Layer Security)란 ?

---

상호 TLS 또는 mTLS는 상호 인증 방법입니다. TLS를 사용하는 mTLS는 인증 및 권한 부여를 모두 수행합니다. mTLS는 트래픽이 클라이언트와 서버 사이의 양방향에서 안전하고 신뢰할 수 있도록 보장합니다.

일반적으로 TLS는 서버 인증과 통신 암호화를 제공합니다. 서버 인증은 클라이언트가 서버의 신원을 확인할 수 있도록 하고, 통신 암호화는 데이터를 암호화하여 중간자 공격으로부터 보호합니다. 그러나 TLS는 기본적으로 클라이언트를 인증하지 않습니다. 이것은 클라이언트가 서버와 통신하는 것을 보호하지만 서버가 클라이언트를 인증하지 않으면 클라이언트가 정말로 해당 서버와 통신하고 있는지 확인할 수 없기 때문에 보안에 취약한 상태가 될 수 있습니다.

mTLS는 이러한 문제를 해결하기 위해 서버와 클라이언트 모두에 대한 인증을 제공합니다. 이를 위해 클라이언트는 자신의 인증서를 서버에 제공하고 서버는 클라이언트의 인증서를 확인하여 클라이언트의 신원을 검증합니다. 이렇게 함으로써 서버는 클라이언트가 실제로 해당 클라이언트인지 확인할 수 있으며 클라이언트는 서버의 신원을 확인할 수 있습니다. 이를 통해 더 높은 보안 수준을 달성할 수 있습니다.

mTLS는 주로 서버 간의 통신, API 호출, 마이크로서비스 아키텍처 등에서 사용됩니다. 그러나 구현 및 관리가 복잡할 수 있으므로 시스템의 요구사항과 상황에 따라 적절하게 사용하는 것이 중요합니다.

### 2. TLS 작동원리

---

[![](https://mermaid.ink/img/pako:eNqNk89u2kAQxl9ltLnEEkiEkjTyIYf-O7W9EPVQcVnw2rFie-liS6EIiUqoB9oDBxygsqVWanJqJERI5EOeCK_foWsbEmOatj7tznzz229mvR3UoApBMmqRDw6xGuSFjjWGzZoF4os-LUL_jvsLfh5EgwCKR0fA-14478uwuycBH8-W8xnwbyM-v0pL0nSizFeLmrK0EixnPTh-XQXuB_zSj2MXd8C_e_yLl3LeUpuAQVQbqPon0hMpHw0HoxWc-5NNcDR1xf5fPcXYyhb2_5wyXTtJrD6g9te9xsZy0L86fHyEgnqwpvLxUCyAuz94P-CDn7C8_Rq5HoS3PT7-BY80W5QeLvDpVq982st45u4omgTRdBQOvXQEyXVzvw_R55vIXYgxuOH1ApY3w2g62TRflLbNowIyCTOxrogfrhPra8g-ISapIVks67glVoVM_B1mOq4bpBULOukBsa5xqjHqWEpat1MqlVZ1cVqllv0Km7rRTtMxw8jlq_rH1aF7h82zTLLJdBOz9nNqUJYKmFbfLVcOoFw5hPJ-ScqILXHzz061jHhHVdWc4Jic2VnFptkEQZlC2BYllnRrVldMDTs2rbatBpJt5pACcpoKttdPFckqNlr30ZeKblN2HyTJ9k36yJO3XkBNbL2ndF3Y_Q2rqdcO?type=png)](https://mermaid.live/edit#pako:eNqNk89u2kAQxl9ltLnEEkiEkjTyIYf-O7W9EPVQcVnw2rFie-liS6EIiUqoB9oDBxygsqVWanJqJERI5EOeCK_foWsbEmOatj7tznzz229mvR3UoApBMmqRDw6xGuSFjjWGzZoF4os-LUL_jvsLfh5EgwCKR0fA-14478uwuycBH8-W8xnwbyM-v0pL0nSizFeLmrK0EixnPTh-XQXuB_zSj2MXd8C_e_yLl3LeUpuAQVQbqPon0hMpHw0HoxWc-5NNcDR1xf5fPcXYyhb2_5wyXTtJrD6g9te9xsZy0L86fHyEgnqwpvLxUCyAuz94P-CDn7C8_Rq5HoS3PT7-BY80W5QeLvDpVq982st45u4omgTRdBQOvXQEyXVzvw_R55vIXYgxuOH1ApY3w2g62TRflLbNowIyCTOxrogfrhPra8g-ISapIVks67glVoVM_B1mOq4bpBULOukBsa5xqjHqWEpat1MqlVZ1cVqllv0Km7rRTtMxw8jlq_rH1aF7h82zTLLJdBOz9nNqUJYKmFbfLVcOoFw5hPJ-ScqILXHzz061jHhHVdWc4Jic2VnFptkEQZlC2BYllnRrVldMDTs2rbatBpJt5pACcpoKttdPFckqNlr30ZeKblN2HyTJ9k36yJO3XkBNbL2ndF3Y_Q2rqdcO)

일반적으로 TLS에서 서버에는 TLS 인증서와 공개/개인 키 쌍이 있지만 클라이언트에는 없습니다. 일반적인 TLS 프로세스는 다음과 같이 작동합니다.

1. 클라이언트가 서버에 연결
2. 서버가 TLS 인증서를 제시
3. 클라이언트는 서버의 인증서를 확인
4. 클라이언트와 서버는 암호화된 TLS 연결을 통해 정보 교환

그러나 mTLS에서는 클라이언트와 서버 모두 인증서를 가지고 있으며 양쪽 모두 공개/개인 키 쌍을 사용하여 인증합니다. 일반 TLS와 비교할 때 mTLS는 양 당사자를 확인하기 위한 추가 단계가 있습니다.

1. 클라이언트가 서버에 연결
2. 서버가 TLS 인증서를 제시
3. 클라이언트는 서버의 인증서를 확인
4. **클라이언트가 TLS 인증서를 제시**
5. **서버는 클라이언트의 인증서를 확인**
6. **서버에서 액세스 권한 부여**
7. 클라이언트와 서버는 암호화된 TLS 연결을 통해 정보를 교환

### 3. 인증서 톺아보기

---

<p align="center">

  <img src="https://user-images.githubusercontent.com/77400522/236472429-6d7b1808-19b5-40b2-b880-7c2c93cae03b.png" />

  <img src="https://user-images.githubusercontent.com/77400522/236472436-476bd074-dab8-4fc7-91b4-e68d76bc44fc.png" />

  <img src="https://user-images.githubusercontent.com/77400522/236472443-829ca8d3-3119-46a4-8df9-5774f1a8a9e6.png" />

<p>

- 여기서 살펴보아야 할 것은 서버와 클라이언트의 `LEAF CERTIFICATE`와 프라이빗 키의 해시값이 동일한지 확인하고 `LEAF CERTIFICATE`의 ISSUER가 ca.crt의 Subject와 동일해야 합니다. 또한 ca.crt는 `ROOT CERTIFICATE`이므로 Subject와 Issuer이 동일합니다.

### 4. 실행 방법

---

#### 4.1 인증서 생성

```bash
cd certificate && go run main.go
```

#### 4.2 서버 실행

```bash
cd server && go run main.go
```

#### 4.3 클라이언트 요청

```bash
cd client
go run main.go -c=a
go run main.go -c=b
```

### 5. 시나리오

---

- 클라이언트 A 요청

```bash
go run main.go -c=a
```

- 서버 커맨드

```bash
2023/05/05 22:45:08 =============== Header ===============
2023/05/05 22:45:08 User-Agent:Go-http-client/1.1
2023/05/05 22:45:08 Accept-Encoding:gzip
2023/05/05 22:45:08 =============== State ===============
2023/05/05 22:45:08 Version: 304
2023/05/05 22:45:08 HandshakeComplete: true
2023/05/05 22:45:08 DidResume: false
2023/05/05 22:45:08 NegotiatedProtocol:
2023/05/05 22:45:08 NegotiatedProtocolIsMutual: true
2023/05/05 22:45:08 Certificate chain:
2023/05/05 22:45:08  0 subject:/C=[Earth]/ST=[Asia]/L=[Mountain]/O=[Client A Company]/OU=[Engineering]/CN=Client A
2023/05/05 22:45:08  issuer:/C=[Earth]/ST=[Asia]/L=[Mountain]/O=[CA Company]/OU=[Engineering]/CN=CA
2023/05/05 22:45:08 =============== End ===============
```

- 클라이언트 B 요청

```bash
go run main.go -c=b
```

- 서버 커맨드

```bash
2023/05/05 22:45:10 =============== Header ===============
2023/05/05 22:45:10 User-Agent:Go-http-client/1.1
2023/05/05 22:45:10 Accept-Encoding:gzip
2023/05/05 22:45:10 =============== State ===============
2023/05/05 22:45:10 Version: 304
2023/05/05 22:45:10 HandshakeComplete: true
2023/05/05 22:45:10 DidResume: false
2023/05/05 22:45:10 NegotiatedProtocol:
2023/05/05 22:45:10 NegotiatedProtocolIsMutual: true
2023/05/05 22:45:10 Certificate chain:
2023/05/05 22:45:10  0 subject:/C=[Earth]/ST=[Asia]/L=[Mountain]/O=[Client B Company]/OU=[Engineering]/CN=Client B
2023/05/05 22:45:10  issuer:/C=[Earth]/ST=[Asia]/L=[Mountain]/O=[CA Company]/OU=[Engineering]/CN=CA
2023/05/05 22:45:10 =============== End ===============
```
