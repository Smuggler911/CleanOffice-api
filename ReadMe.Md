<div class="Header">
<h1>
Clean office API
</h1>
</div>

<h1>https://clean-office-back.ru/</h1>

<div class="auth">
<div class="register" style="font-size: 1.5rem" >
 <h1>
   Регистрация {POST}
</h1>
 <a > 
    http://213.226.127.170:8000/auth/register
</a>
</div>
</div>

```json
{
  "fullname":"",
  "username": "",
  "password": ""
}
```

<div class="Login" style="font-size: 1.5rem" >
 <h1>
   Логин {POST}
</h1>
 <a > 
    http://213.226.127.170:8000/auth/login
</a>
<p>
Возвращает Access Token и Refresh Token
они сразу залетают в куки.
Все операции которые требуют юзер айди проходят через Access Token 
поэтому нужды в получении id юзера нет.Также доступ для редактирования на сайте проверяется через аксес токен

Дополнительно:
 при логине admin и пароле admin 
выдаются права администратора
</p>
</div>


```json
{
  "username": "",
  "password": ""
}
```
<div class="validate" style="font-size: 1.5rem" >
 <h1>
   Профиль Юзера {GET}
</h1>
 <a > 
    http://213.226.127.170:8000/auth/validate
</a>
<p>
Важно нужен  активный  access token в куки
</p>
</div>


<div class="auth" style="font-size: 1.5rem" >
 <h1>
   Добавление Картинки профиля {PUT}
</h1>
 <a > 
    http://213.226.127.170:8000/auth/upload-profile-picture
</a>
<p>
Важно нужен access token активный в куки
</p>
<p>
 имя для  форм даты "picture"
</p>
</div>

<div class="auth" style="font-size: 1.5rem" >
 <h1>
   Изменение профиля {PUT}
</h1>
 <a > 
    http://213.226.127.170:8000/auth/update-user
</a>
<p>
Важно нужен access token активный в куки
</p>
<p>
 имя для  форм даты "picture","username","password"
</p>
</div>

<div class="auth" style="font-size: 1.5rem" >
 <h1>
   Выпуск нового Access Token{POST}
</h1>
 <a > 
    http://213.226.127.170:8000/auth/update-token
</a>
<p>
Важно: нужен access token активный в куки
и refresh token <br />
access token живет примерно 15 минут
нужно обновлять его постоянно <br />
Всего лишь нужно вызывать метод и все  отсальные действия будет выполнять сервер
</p>
</div>

<div class="auth" style="font-size: 1.5rem" >
 <h1>
   Выход из профиля {POST}
</h1>
 <a > 
    http://213.226.127.170:8000/auth/logout
</a>
<p>
Важно: нужен access token активный в куки
</p>
</div>


<div class="auth" style="font-size: 1.5rem" >
 <h1>
   Получение всех пользователей {GET}
</h1>
 <a > 
    http://213.226.127.170:8000/user?limit={сколько надо чтобы показывалось на одной страние}&offset={номер страницы}
</a>
<p>
Важно: нужен access token активный в куки
и обладать правми администратора
</p>
</div>

<div class="auth" style="font-size: 1.5rem" >
 <h1>
    Бан юзера {PUT}
</h1>
 <a > 
    http://213.226.127.170:8000/user/ban-user/:user_id
</a>
<p>
Важно: нужен access token активный в куки
и обладать правми администратора <br/>

</p>
</div>

<div class="auth" style="font-size: 1.5rem" >
 <h1>
    Разбан юзера {PUT}
</h1>
 <a > 
    http://213.226.127.170:8000/user/unban-user/:user_id
</a>
<p>
Важно: нужен access token активный в куки
и обладать правми администратора <br/>

</p>
</div>

<div class="auth" style="font-size: 1.5rem" >
 <h1>
    Просмотр забаненных юзеров {GET}
</h1>
 <a > 
    http://213.226.127.170:8000/user/banned
</a>
<p>
Важно: нужен access token активный в куки
и обладать правми администратора <br/>

</p>
</div>

<div class="offer">
<h1 style="font-size: 4rem">услуги</h1>
</div>

<div class="offer" style="font-size: 1.5rem" >
 <h1>
    Получение списка услуг {GET}
</h1>
 <a > 
    http://213.226.127.170:8000/offer?limit=&offset=

</a>
</div>
<div class="offer" style="font-size: 1.5rem" >
 <h1>
    Получение одной улсуги  {GET}
</h1>
 <a > 
    http://213.226.127.170:8000/offer/:offer_id
</a>
</div>

<div class="offer" style="font-size: 1.5rem" >
 <h1>
    Cоздание услуги {POST}
</h1>
 <a > 
    http://213.226.127.170:8000/offer/create
</a>
<p>
Нужен активный access token и права админа
</p>
<p>
ПОСТ ФОРМА : "picture","name","description"
</p>
</div>

<div class="offer" style="font-size: 1.5rem" >
 <h1>
    Редактирование услуги {PUT}
</h1>
 <a > 
    http://213.226.127.170:8000/offer/update/:offer_id
</a>
<p>
Нужен активный access token и права админа
</p>
<p>
ПОСТ ФОРМА : "picture","name","description"
</p>
</div>

<div class="offer" style="font-size: 1.5rem" >
 <h1>
    Удаление услуги {DELETE}
</h1>
 <a > 
    http://213.226.127.170:8000/offer/delete/:offer_id
</a>
<p>
Нужен активный access token и права админа
</p>

</div>

<div class="offer" style="font-size: 1.5rem" >
 <h1>
    Cоздание ценнового диапазона {POST}
</h1>
 <a > 
    http://213.226.127.170:8000/offer/create-costrange/:offer_id
</a>
<p>
Нужен активный access token и права админа
</p>
<p>
  flat range - это например большие помещения и тд, <br/>
  cost range - это например "oт 10 000"
</p>
</div>

```json
{
  "flatrange": "",
  "costrange": ""
}
```

<div class="offer" style="font-size: 1.5rem" >
 <h1>
    Изменение ценнового диапазона {PUT}
</h1>
 <a > 
    http://213.226.127.170:8000/offer/update-costrange/:offer_id/:cost_id
</a>
<p>
Нужен активный access token и права админа
</p>
<p>
  flat range - это например большие помещения и тд, <br/>
  cost range - это например "oт 10 000"<br/>
  cost_id - это id cost range
</p>
</div>

```json
{
  "flatrange": "",
  "costrange": ""
}
```
<div class="offer" style="font-size: 1.5rem" >
 <h1>
    Удаление ценнового диапазона {DELETE}
</h1>
 <a > 
    http://213.226.127.170:8000/offer/delete-costrange/:offer_id/:cost_id
</a>
<p>
Нужен активный access token и права админа
</p>
<p>
  cost_id - это id cost range
</p>
</div>

<div class="offer" style="font-size: 1.5rem" >
 <h1>
    Показ ценнового диапазона выбранной услуги {DELETE}
</h1>
 <a > 
    http://213.226.127.170:8000/offer/costrange/:offer_id
</a>
<p>
Нужен активный access token и права админа
</p>
</div>

<h1 style="font-size:4rem">Заявки</h1>

<div class="application" style="font-size: 1.5rem" >
 <h1>
    Cоздание звявки {POST}
</h1>
 <a> 
    http://213.226.127.170:8000/application/create/:offer_id/:cost_id
</a>
<p>
Нужен активный access token 
</p>
<p>
Возвращает application_id
для последующего взаимодействия
</p>
</div>

```json
  {
 "time": "",
 "date": ""
}
```

<div class="application" style="font-size: 1.5rem" >
 <h1>
    Отправка кода подтверждения {PUT}
</h1>
 <a> 
    http://213.226.127.170:8000/application/send-verification/:application_Id
</a>
<p>
Нужен активный access token 
</p>

</div>

```json
  {
   "phone": "+7xxxxxxxxxx"
}
```

<div class="application" style="font-size: 1.5rem" >
 <h1>
    Подтверждение {PUT}
</h1>
 <a> 
    http://213.226.127.170:8000/application/verify/:application_Id
</a>
<p>
Нужен активный access token 
</p>

</div>

```json
  {
   "verification": ""
}
```

<div class="application" style="font-size: 1.5rem" >
 <h1>
    Получение забронированных заявок {GET}
</h1>
 <a> 
    http://213.226.127.170:8000/application/booked?limit=&offset=
</a>
<p>
Нужен активный access token и права админа 
</p>
</div>

<div class="application" style="font-size: 1.5rem" >
 <h1>
    Получение поддвержденных заявок {GET}
</h1>
 <a> 
    http://213.226.127.170:8000/application/approved?limit=&offset=
</a>
<p>
Нужен активный access token и права админа 
</p>
</div>

<div class="application" style="font-size: 1.5rem" >
 <h1>
    Получение отклоненных заявок {GET}
</h1>
 <a> 
    http://213.226.127.170:8000/application/declined?limit=&offset=
</a>
<p>
Нужен активный access token и права админа 
</p>
</div>

<div class="application" style="font-size: 1.5rem" >
 <h1>
    Подвердить заявку {PUT}
</h1>
 <a> 
    http://213.226.127.170:8000/application/approve/:application_id
</a>
<p>
Нужен активный access token и права админа 
</p>
</div>

<div class="application" style="font-size: 1.5rem" >
 <h1>
    Отклонить заявку {PUT}
</h1>
 <a> 
    http://213.226.127.170:8000/application/decline/:application_id
</a>
<p>
Нужен активный access token и права админа 
</p>
</div>

<div class="application" style="font-size: 1.5rem" >
 <h1>
    Отметить что выполнено {PUT}
</h1>
 <a> 
    http://213.226.127.170:8000/application/done/:application_id
</a>
<p>
Нужен активный access token и права админа 
</p>
</div>


<h1 style="font-size: 4rem">Отзывы</h1>

<div class="review" style="font-size: 1.5rem" >
 <h1>
    Отсавить отзыв {POST}
</h1>
 <a> 
    http://213.226.127.170:8000/review/create/:application_id
</a>
<p>
Нужен активный access token 
</p>
</div>

```json
{
  "first_rate": 0,
  "second_rate": 0,
  "review": ""
}
```
<div class="review" style="font-size: 1.5rem" >
 <h1>
    Подвтердить{PUT}
</h1>
 <a> 
    http://213.226.127.170:8000/review/approve/:review_id
</a>
<p>
Нужен активный access token и права админа
</p>
</div>

<div class="review" style="font-size: 1.5rem" >
 <h1>
    Отклонить{PUT}
</h1>
 <a> 
    http://213.226.127.170:8000/review/decline/:review_id
</a>
<p>
Нужен активный access token и права админа
</p>
</div>

<div class="review" style="font-size: 1.5rem" >
 <h1>
    Получить комментарии {GET}
</h1>
 <a> 
    http://213.226.127.170:8000/review?limit=&offset=
</a>
</div>

