export function parsearErroresAPI(response: any): string[] {
  const resultado: string[] = [];

  if (response.error) {
    if (typeof response.error === 'string') {
      resultado.push(response.error);

    } else if (Array.isArray(response.error)) {
      response.error.array.forEach((valor: string) => {
        resultado.push(valor)
      });
    }
    else {
      const mapaErrores = response.error.errors;
      const entradas = Object.entries(mapaErrores);
      entradas.forEach((arreglo: any[]) => {
        const campo = arreglo[0];
        arreglo[1].forEach((mensajeError: any) => {
          resultado.push(`${campo}: ${mensajeError}`);
        });
      });
    }
  }

  return resultado;
}
