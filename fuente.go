#include<stdio.h>intmain(){FILE*archivoFuente,*archivoSalida;charc;chardiagonal='';charestrella='*';charespacio='';charsalto='\n';archivoSalida=fopen("salida","w");archivoFuente=fopen("tarea1.cpp","r");while((c=getc(archivoFuente))!=EOF){if(c==espacio||c==salto){putc(espacio,archivoSalida);while((c=getc(archivoFuente))==espacio||c==salto);}if(c==diagonal){c=getc(archivoFuente);if(c==diagonal){while((c=getc(archivoFuente))!=salto);}elseif(c==estrella){boolbandera=true;c=getc(archivoFuente);while(bandera){while(c!=estrella){c=getc(archivoFuente);}if((c=getc(archivoFuente))==diagonal)bandera=false;}}}else{putc(c,archivoSalida);}}return0;}