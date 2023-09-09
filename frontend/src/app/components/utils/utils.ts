export function parsearErroresAPI(response: any): string[] {
  const errores: string[] = [];

  if (response) {
    if (response.error) {
      if (typeof response.error === 'string') {
        // Handle a single error message
        errores.push(response.error);
      } else if (Array.isArray(response.error)) {
        // Handle an array of error messages
        errores.push(...response.error);
      } else if (response.error.errors) {

        const errorObj = response.error.errors;
        for (const campo in errorObj) {
          if (errorObj.hasOwnProperty(campo)) {
            const mensajes = errorObj[campo];
            for (const mensaje of mensajes) {
              errores.push(`${campo}: ${mensaje}`);
            }
          }
        }
//        console.log(errores)
      }
    }
  }

  return errores;
}

