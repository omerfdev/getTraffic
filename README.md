# Trafik Yakalama API

Bu basit Go programı, belirtilen URL'den gelen HTTP GET isteğini yakalar ve yanıtın tamamını geri gönderir.

## Kullanım

- Program, `/capture` endpoint'ına yapılan GET isteklerini dinler.
- Endpoint'a yapılan istekler, bir URL parametresi içermelidir.
- Program, URL'ye yapılan GET isteğinin yanıtını geri döndürür.
- Program, 8080 portunda bir HTTP sunucusu olarak çalışır.

## Dikkat

- Program sadece GET isteklerini kabul eder.
- Gönderilen URL'nin "http://" veya "https://" ile başlaması gereklidir.
- Programı çalıştırmadan önce, gerekli izinleri ve bağımlılıkları sağladığınızdan emin olun.

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Daha fazla bilgi için `LICENSE` dosyasını inceleyebilirsiniz.
