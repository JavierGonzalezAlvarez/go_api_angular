# Front
This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 15

## create project locally
$ npx -p @angular/cli@15.0.0 ng new frontend
* routing: y
* style: scss

https://jwt.io/

## run front
$ npx ng serve

- path: frontend
$ npm install

## create components
path: frontend
* $ npx ng g component components/header
* $ npx ng g component components/footer
* $ npx ng g component components/home
* $ npx ng g component components/profile
* $ npx ng g component components/register
* $ npx ng g component components/modal-invoices

* $ npx ng g component components/header/menu
* $ npx ng g component components/header/title --skip-tests=true

* $ npx ng g component list/listGeneral --skip-tests=true
* $ npx ng g component list/listInvoices --skip-tests=true
* $ npx ng g component list/listInvoicesTable --skip-tests=true
* $ npx ng g component list/listInvoicesTableTotal --skip-tests=true
* $ npx ng g component list/listInvoicesDetails --skip-tests=true

* $ npx ng g component components/seguridad
* $ npx ng g component components/seguridad/login --skip-tests=true
* $ npx ng g component components/seguridad/register --skip-tests=true

## material design
* $ npx ng add @angular/material
* $ npx ng g module material

## service security
* $ npx ng g s components/seguridad/seguridad
* $ npx ng g s components/seguridad/seguridad-interceptor
* $ npx ng g s components/services/modal-invoices

* $ npx ng g s list/service

## interfaces
* $ npx ng g interface interfaces/seguridad
* $ npx ng g interface interfaces/invoices

## authorized user
* $ npx ng g component components/seguridad/autorizado --skip-tests=true

* $ create components/register/seguridad.ts

* $ npx ng g component components/utils/validators
* $ npx ng g component components/seguridad/formulario-autenticacion --skip-tests=true

