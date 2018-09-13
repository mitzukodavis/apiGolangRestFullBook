package models

import "errors"

type ValidationError error

var(
	errorUsername = ValidationError(errors.New("El username no debe estar vacío"))
	errorShortUsername = ValidationError(errors.New("El username es demasiado corto"))
	errorLargeUsername = ValidationError(errors.New("El username es demasiado larho"))

	errorEmail = ValidationError(errors.New("Formo invalido de Email"))

	errorPasswordEncryption = ValidationError(errors.New("no es posible cifrar el texto"))
	errorLogin = ValidationError(errors.New("Usuario o Password incorrecto"))

	)