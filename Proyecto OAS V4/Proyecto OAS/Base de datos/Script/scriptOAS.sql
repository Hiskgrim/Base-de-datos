-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- PostgreSQL version: 9.2
-- Project Site: pgmodeler.com.br
-- Model Author: ---

SET check_function_bodies = false;
-- ddl-end --


-- Database creation must be done outside an multicommand file.
-- These commands were put in this file only for convenience.
-- -- object: new_database | type: DATABASE --
-- CREATE DATABASE new_database
-- ;
-- -- ddl-end --
-- 

-- object: cdve | type: SCHEMA --
CREATE SCHEMA cdve;
-- ddl-end --

-- object: "contratacionDVE" | type: SCHEMA --
CREATE SCHEMA "contratacionDVE";
-- ddl-end --

SET search_path TO pg_catalog,public,cdve,"contratacionDVE";
-- ddl-end --

-- object: public.estado_contrato | type: TABLE --
CREATE TABLE public.estado_contrato(
	id serial NOT NULL,
	nombre character varying NOT NULL,
	fecha_registro date NOT NULL,
	numero_contrato_general integer NOT NULL,
	CONSTRAINT pk_estado_contrato PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.proyecto_curricular | type: TABLE --
CREATE TABLE public.proyecto_curricular(
	id serial NOT NULL,
	nombre character varying(30) NOT NULL,
	descripcion character varying(100),
	id_facultad integer NOT NULL,
	CONSTRAINT pk_proyecto_curricular PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.ciudad | type: TABLE --
CREATE TABLE public.ciudad(
	id serial NOT NULL,
	nombre character varying(20) NOT NULL,
	abreviatura character varying(3),
	descripcion character varying(100),
	id_departamento integer NOT NULL,
	CONSTRAINT pk_ciudad PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.resolucion | type: TABLE --
CREATE TABLE public.resolucion(
	numero int4,
	anio int2,
	descripcion character varying(100),
	entidad character varying(20),
	id_facultad integer NOT NULL,
	CONSTRAINT pk2_resolucion PRIMARY KEY (numero,anio)

);
-- ddl-end --
-- object: public.categoria | type: TABLE --
CREATE TABLE public.categoria(
	id serial NOT NULL,
	nombre character varying(10) NOT NULL,
	descripcion character varying(100),
	CONSTRAINT pk_categoria PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.dedicacion | type: TABLE --
CREATE TABLE public.dedicacion(
	id serial,
	nombre character varying(3),
	descripcion character varying(100),
	CONSTRAINT pk_dedicacion PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.tipo_contrato | type: TABLE --
CREATE TABLE public.tipo_contrato(
	id serial NOT NULL,
	nombre character varying(20),
	descripcion character varying(100),
	CONSTRAINT pk_tipo_contrato PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.docente | type: TABLE --
CREATE TABLE public.docente(
	id serial,
	id_funcionario integer,
	CONSTRAINT pk_docente PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.funcionario | type: TABLE --
CREATE TABLE public.funcionario(
	id serial NOT NULL,
	documento character varying(10) NOT NULL,
	id_ciudad integer NOT NULL,
	telefono_fijo character varying(7),
	telefono_celular character varying(10),
	email character varying(30) NOT NULL,
	primer_nombre character varying(20) NOT NULL,
	segundo_nombre character varying(20),
	primer_apellido character varying NOT NULL,
	segundo_apellido character varying(20),
	sexo char NOT NULL,
	fecha_nacimiento date NOT NULL,
	activo bool NOT NULL,
	rut int4 NOT NULL,
	direccion character varying(30) NOT NULL,
	CONSTRAINT pk_funcionario PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.tipo_contratacion | type: TABLE --
CREATE TABLE public.tipo_contratacion(
	id serial NOT NULL,
	nombre character varying(20) NOT NULL,
	descripcion character varying,
	CONSTRAINT pk_tipo_contratacion PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.facultad | type: TABLE --
CREATE TABLE public.facultad(
	id serial NOT NULL,
	nombre character varying(20) NOT NULL,
	direccion character varying(30),
	CONSTRAINT pk_facultad PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.departamento | type: TABLE --
CREATE TABLE public.departamento(
	id serial NOT NULL,
	nombre character varying(30) NOT NULL,
	abreviatura character varying(3),
	descripcion character varying(100),
	CONSTRAINT pk_departamento PRIMARY KEY (id)

);
-- ddl-end --
-- object: departamento_fk | type: CONSTRAINT --
ALTER TABLE public.ciudad ADD CONSTRAINT departamento_fk FOREIGN KEY (id_departamento)
REFERENCES public.departamento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: funcionario_fk | type: CONSTRAINT --
ALTER TABLE public.docente ADD CONSTRAINT funcionario_fk FOREIGN KEY (id_funcionario)
REFERENCES public.funcionario (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: docente_uq | type: CONSTRAINT --
ALTER TABLE public.docente ADD CONSTRAINT docente_uq UNIQUE (id_funcionario);
-- ddl-end --


-- object: facultad_fk | type: CONSTRAINT --
ALTER TABLE public.proyecto_curricular ADD CONSTRAINT facultad_fk FOREIGN KEY (id_facultad)
REFERENCES public.facultad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: ciudad_fk | type: CONSTRAINT --
ALTER TABLE public.funcionario ADD CONSTRAINT ciudad_fk FOREIGN KEY (id_ciudad)
REFERENCES public.ciudad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: facultad_fk | type: CONSTRAINT --
ALTER TABLE public.resolucion ADD CONSTRAINT facultad_fk FOREIGN KEY (id_facultad)
REFERENCES public.facultad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: public.clasificacion | type: TABLE --
CREATE TABLE public.clasificacion(
	id serial,
	fecha_registro date,
	id_dedicacion integer NOT NULL,
	id_categoria integer NOT NULL,
	id_docente integer NOT NULL,
	CONSTRAINT pk_clasificacion PRIMARY KEY (id)

);
-- ddl-end --
-- object: dedicacion_fk | type: CONSTRAINT --
ALTER TABLE public.clasificacion ADD CONSTRAINT dedicacion_fk FOREIGN KEY (id_dedicacion)
REFERENCES public.dedicacion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: categoria_fk | type: CONSTRAINT --
ALTER TABLE public.clasificacion ADD CONSTRAINT categoria_fk FOREIGN KEY (id_categoria)
REFERENCES public.categoria (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: docente_fk | type: CONSTRAINT --
ALTER TABLE public.clasificacion ADD CONSTRAINT docente_fk FOREIGN KEY (id_docente)
REFERENCES public.docente (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: public.carga_academica | type: TABLE --
CREATE TABLE public.carga_academica(
	id serial,
	horas_semanales int2,
	id_proyecto_curricular integer NOT NULL,
	id_docente integer NOT NULL,
	CONSTRAINT pk_carga_academica PRIMARY KEY (id)

);
-- ddl-end --
-- object: proyecto_curricular_fk | type: CONSTRAINT --
ALTER TABLE public.carga_academica ADD CONSTRAINT proyecto_curricular_fk FOREIGN KEY (id_proyecto_curricular)
REFERENCES public.proyecto_curricular (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: docente_fk | type: CONSTRAINT --
ALTER TABLE public.carga_academica ADD CONSTRAINT docente_fk FOREIGN KEY (id_docente)
REFERENCES public.docente (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: public.cancelacion_contrato | type: TABLE --
CREATE TABLE public.cancelacion_contrato(
	id serial NOT NULL,
	fecha_registro date,
	descripcion character varying(200),
	numero_contrato_general integer,
	CONSTRAINT pk_cancelacion_contrato PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.contrato_general | type: TABLE --
CREATE TABLE public.contrato_general(
	numero serial,
	vigencia integer,
	objeto_contrato character varying(9000),
	"plazo_ejecución" integer,
	clausula_registro_presupuestal boolean,
	sede_solicitante character varying(50),
	dependencia_solicitante character varying(10),
	numero_solicitud_necesidad integer,
	numero_cdp smallint,
	contratista numeric(16,0),
	valor_contrato numeric(16,2),
	justificacion character varying,
	descripcion_forma_pago character varying,
	condiciones character varying,
	fecha_registro date,
	tipologia_contrato integer,
	tipo_compromiso integer,
	modalidad_seleccion integer,
	procedimiento integer,
	regimen_contratacion integer,
	tipo_gasto integer,
	toma_gasto_inversion integer,
	origen_presupuesto integer,
	origen_recursos integer,
	tipo_moneda integer,
	valor_contrato_me numeric(16,3),
	valor_tasa_cambio numeric(15,10),
	tipo_control integer,
	observaciones character varying(1000),
	clase_contratista integer,
	convenio character varying,
	numero_constancia integer,
	estado boolean,
	registro_presupuestal integer,
	id_tipo_contrato integer NOT NULL,
	CONSTRAINT pk_contrato_general PRIMARY KEY (numero)

);
-- ddl-end --
-- object: public.contrato_especial | type: TABLE --
CREATE TABLE public.contrato_especial(
	id serial,
	fecha_inicio date,
	vigencia int2,
	numero_contrato_general integer NOT NULL,
	id_funcionario integer NOT NULL,
	numero_resolucion int4,
	anio_resolucion int2,
	CONSTRAINT pk_contrato_especial PRIMARY KEY (id)

);
-- ddl-end --
-- object: contrato_general_fk | type: CONSTRAINT --
ALTER TABLE public.contrato_especial ADD CONSTRAINT contrato_general_fk FOREIGN KEY (numero_contrato_general)
REFERENCES public.contrato_general (numero) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: contrato_general_fk | type: CONSTRAINT --
ALTER TABLE public.cancelacion_contrato ADD CONSTRAINT contrato_general_fk FOREIGN KEY (numero_contrato_general)
REFERENCES public.contrato_general (numero) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: cancelacion_contrato_uq | type: CONSTRAINT --
ALTER TABLE public.cancelacion_contrato ADD CONSTRAINT cancelacion_contrato_uq UNIQUE (numero_contrato_general);
-- ddl-end --


-- object: contrato_general_fk | type: CONSTRAINT --
ALTER TABLE public.estado_contrato ADD CONSTRAINT contrato_general_fk FOREIGN KEY (numero_contrato_general)
REFERENCES public.contrato_general (numero) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: funcionario_fk | type: CONSTRAINT --
ALTER TABLE public.contrato_especial ADD CONSTRAINT funcionario_fk FOREIGN KEY (id_funcionario)
REFERENCES public.funcionario (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: tipo_contrato_fk | type: CONSTRAINT --
ALTER TABLE public.contrato_general ADD CONSTRAINT tipo_contrato_fk FOREIGN KEY (id_tipo_contrato)
REFERENCES public.tipo_contrato (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: public.solicitud_necesidad | type: TABLE --
CREATE TABLE public.solicitud_necesidad(
	id serial,
	numero int2,
	dependencia int2,
	dependencia_destino int2,
	ordenador_gasto int2,
	unidad_ejecutora int2,
	justificacion character varying(1000),
	valor int4,
	fecha_evaluacion date,
	justificacion_rechazo character varying(1000),
	id_tipo_contratacion integer,
	CONSTRAINT pk_solicitud_necesidad PRIMARY KEY (id)

);
-- ddl-end --
-- object: tipo_contratacion_fk | type: CONSTRAINT --
ALTER TABLE public.solicitud_necesidad ADD CONSTRAINT tipo_contratacion_fk FOREIGN KEY (id_tipo_contratacion)
REFERENCES public.tipo_contratacion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: resolucion_fk | type: CONSTRAINT --
ALTER TABLE public.contrato_especial ADD CONSTRAINT resolucion_fk FOREIGN KEY (numero_resolucion,anio_resolucion)
REFERENCES public.resolucion (numero,anio) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: public.asignatura | type: TABLE --
CREATE TABLE public.asignatura(
	id serial,
	nombre character varying(30) NOT NULL,
	descripcion character varying(1000),
	CONSTRAINT pk_asignatura PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.grupo | type: TABLE --
CREATE TABLE public.grupo(
	id serial,
	horario character varying NOT NULL,
	horas_semanales int2 NOT NULL,
	id_asignatura integer NOT NULL,
	CONSTRAINT pk_grupo PRIMARY KEY (id)

);
-- ddl-end --
-- object: asignatura_fk | type: CONSTRAINT --
ALTER TABLE public.grupo ADD CONSTRAINT asignatura_fk FOREIGN KEY (id_asignatura)
REFERENCES public.asignatura (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: public.docente_grupo | type: TABLE --
CREATE TABLE public.docente_grupo(
	id serial,
	horas int2 NOT NULL,
	id_grupo integer NOT NULL,
	id_docente integer NOT NULL,
	CONSTRAINT pk_docente_grupo PRIMARY KEY (id)

);
-- ddl-end --
-- object: grupo_fk | type: CONSTRAINT --
ALTER TABLE public.docente_grupo ADD CONSTRAINT grupo_fk FOREIGN KEY (id_grupo)
REFERENCES public.grupo (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: docente_fk | type: CONSTRAINT --
ALTER TABLE public.docente_grupo ADD CONSTRAINT docente_fk FOREIGN KEY (id_docente)
REFERENCES public.docente (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --



