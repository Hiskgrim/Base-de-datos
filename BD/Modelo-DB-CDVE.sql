-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- PostgreSQL version: 9.1
-- Project Site: pgmodeler.com.br
-- Model Author: ---

SET check_function_bodies = false;
-- ddl-end --


-- Database creation must be done outside an multicommand file.
-- These commands were put in this file only for convenience.
-- -- object: udistrital | type: DATABASE --
-- CREATE DATABASE udistrital
-- 	ENCODING = 'UTF8'
-- 	LC_COLLATE = 'Spanish_Mexico.UTF8'
-- 	LC_CTYPE = 'Spanish_Mexico.UTF8'
-- 	TABLESPACE = pg_default
-- 	OWNER = postgres
-- ;
-- -- ddl-end --
-- 

-- object: core | type: SCHEMA --
CREATE SCHEMA core;
ALTER SCHEMA core OWNER TO postgres;
-- ddl-end --

-- object: kronos | type: SCHEMA --
CREATE SCHEMA kronos;
ALTER SCHEMA kronos OWNER TO postgres;
COMMENT ON SCHEMA kronos IS 'esquema sistema financiero';
-- ddl-end --
-- ddl-end --

-- object: argo | type: SCHEMA --
CREATE SCHEMA argo;
ALTER SCHEMA argo OWNER TO postgres;
-- ddl-end --

-- object: agora | type: SCHEMA --
CREATE SCHEMA agora;
ALTER SCHEMA agora OWNER TO postgres;
-- ddl-end --

-- object: personal | type: SCHEMA --
CREATE SCHEMA personal;
ALTER SCHEMA personal OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,core,kronos,argo,agora,personal;
-- ddl-end --

-- object: core.parametro_entidad | type: TABLE --
CREATE TABLE core.parametro_entidad(
	id integer NOT NULL DEFAULT nextval('core.parametro_entidad_id_seq'::regclass),
	nit character varying(10),
	nombre character varying(30) NOT NULL,
	fecha_creacion date NOT NULL,
	id_tipo_entidad integer NOT NULL,
	id_estado integer NOT NULL,
	fecha_modificacion date DEFAULT now(),
	CONSTRAINT id_entidad PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON CONSTRAINT id_entidad ON core.parametro_entidad IS 'id de la entidad';
-- ddl-end --
ALTER TABLE core.parametro_entidad OWNER TO postgres;
-- ddl-end --

-- object: core.banco | type: TABLE --
CREATE TABLE core.banco(
	id_codigo numeric(5,0) NOT NULL,
	nombre_banco character varying(150) NOT NULL,
	denominacion_banco character varying(250),
	descripcion character varying(150),
	nit character varying(10) NOT NULL,
	codigo_superintendencia numeric(3,0) NOT NULL,
	codigo_ach numeric(3,0) NOT NULL,
	estado_activo boolean NOT NULL DEFAULT true,
	CONSTRAINT bancos_pkey PRIMARY KEY (id_codigo),
	CONSTRAINT nit_uk UNIQUE (nit),
	CONSTRAINT codigo_ach_uk UNIQUE (codigo_ach),
	CONSTRAINT codigo_superintendencia_uk UNIQUE (codigo_superintendencia)

);
-- ddl-end --
COMMENT ON TABLE core.banco IS 'Tabla que parametriza todas las Entidades Bancarias en Colombia';
-- ddl-end --
COMMENT ON COLUMN core.banco.id_codigo IS 'Identificación Unica de las Entidades Bancarias';
-- ddl-end --
COMMENT ON COLUMN core.banco.nombre_banco IS 'Nombre del Banco que se Parametriza';
-- ddl-end --
COMMENT ON COLUMN core.banco.denominacion_banco IS 'Nombre Legal de la Entidad Bancaria para mayor referencia';
-- ddl-end --
COMMENT ON COLUMN core.banco.descripcion IS 'campo en el que se puede registrar una descripcion para el banco';
-- ddl-end --
COMMENT ON COLUMN core.banco.nit IS 'campo en el que se debe registrar el nit del banco';
-- ddl-end --
COMMENT ON COLUMN core.banco.codigo_superintendencia IS 'campo en el que se debe registrar el codigo de la superintendencia ';
-- ddl-end --
COMMENT ON COLUMN core.banco.codigo_ach IS 'campo en el que se debe registrar el codigo ACH';
-- ddl-end --
COMMENT ON COLUMN core.banco.estado_activo IS 'campo en el que se indica si el banco se encuentra activo o no';
-- ddl-end --
ALTER TABLE core.banco OWNER TO postgres;
-- ddl-end --

-- object: core.ciudad | type: TABLE --
CREATE TABLE core.ciudad(
	id_ciudad numeric(10,0) NOT NULL DEFAULT nextval('core.ciudad_id_ciudad_seq'::regclass),
	id_departamento integer NOT NULL,
	nombre character varying(60) NOT NULL,
	abreviatura character varying(3),
	descripcion character varying(300),
	estado character varying(10) NOT NULL,
	ab_pais character varying(4),
	departamento character varying(50),
	poblacion integer,
	longitud numeric(10,3),
	latitud numeric(10,3),
	CONSTRAINT ckc_estado_ciudad CHECK (((estado)::text = ANY (ARRAY[('Activo'::character varying)::text, ('Inactivo'::character varying)::text]))),
	CONSTRAINT pk_ciudad PRIMARY KEY (id_ciudad)

);
-- ddl-end --
-- object: ciudad_pk | type: INDEX --
CREATE UNIQUE INDEX ciudad_pk ON core.ciudad
	USING btree
	(
	  id_ciudad
	)	WITH (FILLFACTOR = 90);
-- ddl-end --

-- object: ciudadxdepto_fk | type: INDEX --
CREATE INDEX ciudadxdepto_fk ON core.ciudad
	USING btree
	(
	  id_departamento
	)	WITH (FILLFACTOR = 90);
-- ddl-end --


COMMENT ON TABLE core.ciudad IS 'Tabla que Parametriza la gran mayoria de todas las Ciudades Mundiales';
-- ddl-end --
ALTER TABLE core.ciudad OWNER TO postgres;
-- ddl-end --

-- object: core.departamento | type: TABLE --
CREATE TABLE core.departamento(
	id_departamento numeric(10,0) NOT NULL DEFAULT nextval('core.departamento_id_departamento_seq'::regclass),
	id_pais integer NOT NULL,
	nombre character varying(60) NOT NULL,
	abreviatura character varying(3),
	descripcion character varying(300),
	estado character varying(10) NOT NULL,
	ab_pais character varying(4),
	poblacion integer,
	area integer,
	capital_departamento character varying(50),
	departamento_cap_pais character varying(50),
	CONSTRAINT ckc_estado_departam CHECK (((estado)::text = ANY (ARRAY[('Activo'::character varying)::text, ('Inactivo'::character varying)::text]))),
	CONSTRAINT pk_departamento PRIMARY KEY (id_departamento)

);
-- ddl-end --
-- object: departamento_pk | type: INDEX --
CREATE UNIQUE INDEX departamento_pk ON core.departamento
	USING btree
	(
	  id_departamento
	)	WITH (FILLFACTOR = 90);
-- ddl-end --

-- object: relationship_1_fk | type: INDEX --
CREATE INDEX relationship_1_fk ON core.departamento
	USING btree
	(
	  id_pais
	)	WITH (FILLFACTOR = 90);
-- ddl-end --


COMMENT ON TABLE core.departamento IS 'Tabla que Parametriza la gran mayoria de todos las Departamentos/Estados Mundiales';
-- ddl-end --
ALTER TABLE core.departamento OWNER TO postgres;
-- ddl-end --

-- object: core.pais | type: TABLE --
CREATE TABLE core.pais(
	id_pais numeric(10,0) NOT NULL DEFAULT nextval('core.pais_id_pais_seq'::regclass),
	nombre_pais character varying(50) NOT NULL,
	abreviatura character varying(3),
	estado character varying(10) NOT NULL,
	capital_pais character varying(50),
	provincia_pais character varying(50),
	area_pais integer,
	poblacion_pais integer,
	codigo_pais integer,
	CONSTRAINT ckc_estado_pais CHECK (((estado)::text = ANY (ARRAY[('Activo'::character varying)::text, ('Inactivo'::character varying)::text]))),
	CONSTRAINT pk_pais PRIMARY KEY (id_pais)

);
-- ddl-end --
-- object: pais_pk | type: INDEX --
CREATE UNIQUE INDEX pais_pk ON core.pais
	USING btree
	(
	  id_pais
	)	WITH (FILLFACTOR = 90);
-- ddl-end --


COMMENT ON TABLE core.pais IS 'Tabla que Parametriza todos los Paises del Mundo';
-- ddl-end --
ALTER TABLE core.pais OWNER TO postgres;
-- ddl-end --

-- object: core.snies_nucleo_basico | type: TABLE --
CREATE TABLE core.snies_nucleo_basico(
	id_area numeric(10,0) NOT NULL,
	id_nucleo numeric(10,0) NOT NULL,
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	estado character varying(15) NOT NULL,
	CONSTRAINT ckc_estado_nucleo CHECK (((estado)::text = ANY (ARRAY[('ACTIVO'::character varying)::text, ('INACTIVO'::character varying)::text]))),
	CONSTRAINT id_nucleo PRIMARY KEY (id_nucleo)

);
-- ddl-end --
COMMENT ON TABLE core.snies_nucleo_basico IS 'Tabla que parametriza el Nucleo Básico del Conocimiento NBC para las Profesiones segun el SNIES';
-- ddl-end --
COMMENT ON COLUMN core.snies_nucleo_basico.id_area IS 'Codigo Manejado por el SNIES para el Area (FK) de la tabla snies_area';
-- ddl-end --
COMMENT ON COLUMN core.snies_nucleo_basico.id_nucleo IS 'Codigo Manejado por el SNIES para el NBC';
-- ddl-end --
COMMENT ON COLUMN core.snies_nucleo_basico.nombre IS 'Nombre del Nucleo Basico del Conocimiento';
-- ddl-end --
COMMENT ON COLUMN core.snies_nucleo_basico.descripcion IS 'Descripcion del NBC si se dispone';
-- ddl-end --
COMMENT ON COLUMN core.snies_nucleo_basico.estado IS 'Estado Activo o Inactivo del Nucleo Basico del Conocimiento';
-- ddl-end --
ALTER TABLE core.snies_nucleo_basico OWNER TO postgres;
-- ddl-end --

-- object: agora.informacion_persona_natural | type: TABLE --
CREATE TABLE agora.informacion_persona_natural(
	tipo_documento numeric(5,0) NOT NULL,
	num_documento_persona numeric(15,0) NOT NULL,
	digito_verificacion numeric(2,0) NOT NULL,
	primer_apellido character varying(50) NOT NULL,
	segundo_apellido character varying(50),
	primer_nombre character varying(50) NOT NULL,
	segundo_nombre character varying(50),
	cargo character varying(150),
	id_pais_nacimiento numeric(5,0) NOT NULL,
	perfil numeric(5,0) NOT NULL,
	profesion character varying(150),
	especialidad character varying(100),
	monto_capital_autorizado numeric(16,2),
	genero character varying(20),
	grupo_etnico character varying(20),
	comunidad_lgbt boolean NOT NULL DEFAULT false,
	cabeza_familia boolean NOT NULL DEFAULT false,
	personas_a_cargo boolean NOT NULL DEFAULT false,
	numero_personas_a_cargo numeric(10,0),
	estado_civil character varying(20) NOT NULL DEFAULT 'SOLTERO'::character varying,
	discapacitado boolean NOT NULL DEFAULT false,
	tipo_discapacidad character varying(20),
	declarante_renta boolean NOT NULL DEFAULT false,
	medicina_prepagada boolean NOT NULL DEFAULT false,
	valor_uvt_prepagada numeric(10,4),
	cuenta_ahorro_afc boolean NOT NULL DEFAULT false,
	num_cuenta_bancaria_afc character varying(15),
	id_entidad_bancaria_afc numeric(5,0),
	interes_vivienda_afc numeric(10,4),
	dependiente_hijo_menor_edad boolean NOT NULL DEFAULT false,
	dependiente_hijo_menos23_estudiando boolean NOT NULL DEFAULT false,
	dependiente_hijo_mas23_discapacitado boolean NOT NULL DEFAULT false,
	dependiente_conyuge boolean NOT NULL DEFAULT false,
	dependiente_padre_o_hermano boolean NOT NULL DEFAULT false,
	id_nucleo_basico numeric(10,0),
	id_arl integer,
	id_eps integer,
	id_fondo_pension integer,
	id_caja_compensacion integer,
	id_nit_arl numeric(20,0),
	id_nit_eps numeric(20,0),
	id_nit_fondo_pension numeric(20,0),
	id_nit_caja_compensacion numeric(20,0),
	fecha_expedicion_documento date NOT NULL DEFAULT '1900-01-01'::date,
	id_ciudad_expedicion_documento numeric(5,0) NOT NULL DEFAULT 96,
	CONSTRAINT ckc_estado_civil CHECK (((estado_civil)::text = ANY (ARRAY[('SOLTERO'::character varying)::text, ('CASADO'::character varying)::text, ('UNION LIBRE'::character varying)::text, ('VIUDO'::character varying)::text, ('DIVORCIADO'::character varying)::text]))),
	CONSTRAINT ckc_etnia CHECK (((grupo_etnico)::text = ANY (ARRAY[('AFRODESCENDIENTES'::character varying)::text, ('INDIGENAS'::character varying)::text, ('RAIZALES'::character varying)::text, ('ROM'::character varying)::text]))),
	CONSTRAINT ckc_genero CHECK (((genero)::text = ANY (ARRAY[('MASCULINO'::character varying)::text, ('FEMENINO'::character varying)::text, ('NO APLICA'::character varying)::text]))),
	CONSTRAINT ckc_tipo_discapacidad CHECK (((tipo_discapacidad)::text = ANY (ARRAY[('FISICA'::character varying)::text, ('SENSORIAL'::character varying)::text, ('AUDITIVA'::character varying)::text, ('VISUAL'::character varying)::text, ('PSIQUICA'::character varying)::text, ('MENTAL'::character varying)::text]))),
	CONSTRAINT "pki_idPersona" PRIMARY KEY (num_documento_persona)

);
-- ddl-end --
-- object: fki_id_pais_nacimiento | type: INDEX --
CREATE INDEX fki_id_pais_nacimiento ON agora.informacion_persona_natural
	USING btree
	(
	  id_pais_nacimiento
	)	WITH (FILLFACTOR = 90);
-- ddl-end --

-- object: fki_id_perfil | type: INDEX --
CREATE INDEX fki_id_perfil ON agora.informacion_persona_natural
	USING btree
	(
	  perfil
	)	WITH (FILLFACTOR = 90);
-- ddl-end --

-- object: "idParametroTipoDoc" | type: INDEX --
CREATE INDEX "idParametroTipoDoc" ON agora.informacion_persona_natural
	USING btree
	(
	  tipo_documento
	)	WITH (FILLFACTOR = 90);
-- ddl-end --


ALTER TABLE agora.informacion_persona_natural OWNER TO postgres;
-- ddl-end --

-- object: agora.parametro_estandar | type: TABLE --
CREATE TABLE agora.parametro_estandar(
	id_parametro numeric(10,0) NOT NULL,
	clase_parametro character varying(30),
	valor_parametro character varying(80) NOT NULL,
	descripcion_parametro character varying(50),
	CONSTRAINT "pk_idParametro" PRIMARY KEY (id_parametro)

);
-- ddl-end --
ALTER TABLE agora.parametro_estandar OWNER TO postgres;
-- ddl-end --

-- object: kronos.unidad_ejecutora | type: TABLE --
CREATE TABLE kronos.unidad_ejecutora(
	id integer NOT NULL,
	nombre character varying(300),
	descripcion character varying(500),
	CONSTRAINT pk_unidad_ejecutora PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE kronos.unidad_ejecutora IS 'tabla que almacena las unidades ejecutoras involucradas en el sistema';
-- ddl-end --
COMMENT ON COLUMN kronos.unidad_ejecutora.id IS 'identificador unico de la unidad ejecutora';
-- ddl-end --
COMMENT ON COLUMN kronos.unidad_ejecutora.nombre IS 'nombre de la unidad ejecutora';
-- ddl-end --
COMMENT ON COLUMN kronos.unidad_ejecutora.descripcion IS 'descripcion de la unidad ejecutora ';
-- ddl-end --
ALTER TABLE kronos.unidad_ejecutora OWNER TO postgres;
-- ddl-end --

-- object: argo.tipo_contrato | type: TABLE --
CREATE TABLE argo.tipo_contrato(
	id integer NOT NULL DEFAULT nextval('argo.tipo_contrato_id_seq1'::regclass),
	tipo_contrato character varying,
	id_grupo_tipo_contrato integer,
	estado boolean DEFAULT true,
	CONSTRAINT pk_tipo_contrato PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE argo.tipo_contrato IS 'tipo especifico del contrato ';
-- ddl-end --
COMMENT ON COLUMN argo.tipo_contrato.id IS 'identificador unico del registro';
-- ddl-end --
COMMENT ON COLUMN argo.tipo_contrato.tipo_contrato IS 'descripcion del tipo de contrato';
-- ddl-end --
COMMENT ON COLUMN argo.tipo_contrato.id_grupo_tipo_contrato IS 'identificador del grupo_tipo_contrato al cual pertene el tipo_contrato';
-- ddl-end --
COMMENT ON COLUMN argo.tipo_contrato.estado IS 'estado del tipo de contrato';
-- ddl-end --
ALTER TABLE argo.tipo_contrato OWNER TO postgres;
-- ddl-end --

-- object: argo.supervisor_contrato | type: TABLE --
CREATE TABLE argo.supervisor_contrato(
	id integer NOT NULL DEFAULT nextval('argo.supervisor_contrato_id_seq'::regclass),
	nombre character varying NOT NULL,
	documento integer NOT NULL,
	cargo character varying NOT NULL,
	sede_supervisor character varying,
	dependencia_supervisor character varying,
	tipo integer,
	estado boolean DEFAULT true,
	digito_verificacion integer,
	CONSTRAINT pk_supervisor_contrato PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE argo.supervisor_contrato IS 'tabla que almacena la informacion del supervisor del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.supervisor_contrato.id IS 'identificador unico de la tabla ';
-- ddl-end --
COMMENT ON COLUMN argo.supervisor_contrato.nombre IS 'Nombre del supervisor ';
-- ddl-end --
COMMENT ON COLUMN argo.supervisor_contrato.documento IS 'documento de identificacion del supervisor';
-- ddl-end --
COMMENT ON COLUMN argo.supervisor_contrato.cargo IS 'Cargo del supervisor';
-- ddl-end --
COMMENT ON COLUMN argo.supervisor_contrato.tipo IS 'campo tipo supervisor se crea con el fin de poder tipificar a futuro los supervisores que son funcionarios y lo que no los son';
-- ddl-end --
COMMENT ON COLUMN argo.supervisor_contrato.estado IS 'estado del registro';
-- ddl-end --
COMMENT ON COLUMN argo.supervisor_contrato.digito_verificacion IS 'Digito de Verificacion supervisor';
-- ddl-end --
ALTER TABLE argo.supervisor_contrato OWNER TO postgres;
-- ddl-end --

-- object: argo.parametros | type: TABLE --
CREATE TABLE argo.parametros(
	id_parametro integer NOT NULL DEFAULT nextval('argo.parametros_id_parametro_seq'::regclass),
	descripcion character varying NOT NULL,
	codigo_contraloria character varying,
	rel_parametro integer NOT NULL,
	estado_registro boolean NOT NULL DEFAULT true,
	fecha_registro date NOT NULL,
	CONSTRAINT pk_id_parametros_pr PRIMARY KEY (id_parametro)

);
-- ddl-end --
ALTER TABLE argo.parametros OWNER TO postgres;
-- ddl-end --

-- object: argo.contrato_general | type: TABLE --
CREATE TABLE argo.contrato_general(
	numero_contrato character varying NOT NULL DEFAULT nextval('argo.numero_unico_contrato_seq'::regclass),
	vigencia integer NOT NULL,
	objeto_contrato character varying(9000),
	plazo_ejecucion integer NOT NULL,
	forma_pago integer NOT NULL,
	ordenador_gasto character varying(14),
	clausula_registro_presupuestal boolean,
	sede_solicitante character varying(50),
	dependencia_solicitante character varying(10),
	numero_solicitud_necesidad integer NOT NULL,
	numero_cdp integer NOT NULL,
	contratista numeric(16,0) NOT NULL,
	unidad_ejecucion integer NOT NULL,
	valor_contrato numeric(16,2) NOT NULL,
	justificacion character varying NOT NULL,
	descripcion_forma_pago character varying NOT NULL,
	condiciones character varying NOT NULL,
	unidad_ejecutora integer NOT NULL,
	fecha_registro date NOT NULL DEFAULT ('now'::text)::date,
	tipologia_contrato integer NOT NULL,
	tipo_compromiso integer NOT NULL,
	modalidad_seleccion integer NOT NULL,
	procedimiento integer NOT NULL,
	regimen_contratacion integer NOT NULL,
	tipo_gasto integer NOT NULL,
	tema_gasto_inversion integer NOT NULL,
	origen_presupueso integer NOT NULL,
	origen_recursos integer NOT NULL,
	tipo_moneda integer NOT NULL,
	valor_contrato_me numeric(16,3),
	valor_tasa_cambio numeric(15,10),
	tipo_control integer,
	observaciones character varying(1000),
	supervisor integer,
	clase_contratista integer NOT NULL,
	convenio character varying,
	numero_constancia integer,
	estado boolean DEFAULT true,
	resgistro_presupuestal integer,
	tipo_contrato integer NOT NULL,
	lugar_ejecucion bigint NOT NULL,
	actividades character varying,
	CONSTRAINT pk_contrato_general PRIMARY KEY (numero_contrato,vigencia),
	CONSTRAINT ck_numero_contrato_contrato_general CHECK (((numero_contrato)::text ~ '^((FPL|DVE|DPL|RES)[0-9]+|[0-9]+)$'::text))

);
-- ddl-end --
-- object: fki_contrato_general_ordenador | type: INDEX --
CREATE INDEX fki_contrato_general_ordenador ON argo.contrato_general
	USING btree
	(
	  ordenador_gasto
	)	WITH (FILLFACTOR = 90);
-- ddl-end --

-- object: fki_funcionario_ordenador | type: INDEX --
CREATE INDEX fki_funcionario_ordenador ON argo.contrato_general
	USING btree
	(
	  ordenador_gasto
	)	WITH (FILLFACTOR = 90);
-- ddl-end --

-- object: fki_lugar_ejecucion_contrato_general | type: INDEX --
CREATE INDEX fki_lugar_ejecucion_contrato_general ON argo.contrato_general
	USING btree
	(
	  lugar_ejecucion
	)	WITH (FILLFACTOR = 90);
-- ddl-end --

-- object: fki_parametro_unidad_tiempo_ejecucion_parametros | type: INDEX --
CREATE INDEX fki_parametro_unidad_tiempo_ejecucion_parametros ON argo.contrato_general
	USING btree
	(
	  unidad_ejecucion
	)	WITH (FILLFACTOR = 90);
-- ddl-end --

-- object: fki_supervisor_contrato_contrato_general | type: INDEX --
CREATE INDEX fki_supervisor_contrato_contrato_general ON argo.contrato_general
	USING btree
	(
	  supervisor
	)	WITH (FILLFACTOR = 90);
-- ddl-end --

-- object: "fki_tipo_contrato_contrato_generaL" | type: INDEX --
CREATE INDEX "fki_tipo_contrato_contrato_generaL" ON argo.contrato_general
	USING btree
	(
	  tipo_contrato
	)	WITH (FILLFACTOR = 90);
-- ddl-end --


COMMENT ON COLUMN argo.contrato_general.numero_contrato IS 'identificador unico del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.vigencia IS 'vigencia del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.objeto_contrato IS 'descripcion objeto del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.plazo_ejecucion IS 'plazo de ejecucion del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.forma_pago IS 'forma de pago del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.ordenador_gasto IS 'identificador foraneo del ordenador del gasto del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.clausula_registro_presupuestal IS 'clausula de presupuesto que indica si el contrato esta sujeto a disponibilidad de ddinero';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.sede_solicitante IS 'sede solicitante del contrato, en la solcitud de necesidad se especifica la sede que realiza la solicitud.';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.dependencia_solicitante IS 'dependencia o convenio solicitante';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.numero_solicitud_necesidad IS 'numero de la solicitud asociada al contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.numero_cdp IS 'numero del cdp con el cual se registra el contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.contratista IS 'Identificacion del Contratista';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.unidad_ejecucion IS 'unidad ejecucion bajo la cual se ejecuta el contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.valor_contrato IS 'corresponde al valor del contrato o de una orden ';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.justificacion IS 'justificacion del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.descripcion_forma_pago IS 'descripcion de la forma de pago del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.condiciones IS 'condiciones del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.unidad_ejecutora IS 'identificador foraneo de la unidad ejecutora a cargo del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.fecha_registro IS 'fecha de registro del contrato general';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.tipologia_contrato IS 'Referencia Foranea al Tipo de Contrato que se esta creando';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.tipo_compromiso IS 'referencia foranea al tipo de compromiso';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.modalidad_seleccion IS 'identificador foraneo al parameto modalidad de seleccion ';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.procedimiento IS 'identificador foraneo del parametro procedimiento de contratacion';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.regimen_contratacion IS 'identificador foraneo al parametro regimen de contratacion';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.tipo_gasto IS 'identificador foraneo al parametro tipo de gasto ';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.tema_gasto_inversion IS 'identificador foraneo al parametro tema a que corresponde el gasto o la inversion';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.origen_presupueso IS 'identificador foraneo al parametro origen del presupuesto';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.origen_recursos IS 'identificador foraneo al parametro origen de los recursos';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.tipo_moneda IS 'identificador foraneo al parametro tipo de moneda ';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.valor_contrato_me IS 'valor del contrato en moneda extrangera';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.valor_tasa_cambio IS 'Valor tasa de cambio en el momento de la suscripcion del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.tipo_control IS 'identificador foraneo al parametro tipo de control ';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.observaciones IS 'observaciones del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.supervisor IS 'identificador foraneo del supervisor del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.convenio IS 'convenio asociado al contrato en caso de existir';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.numero_constancia IS 'numero de constancia del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.estado IS 'estado del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.resgistro_presupuestal IS 'identificador del registro presupuestal asociado';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.tipo_contrato IS 'identificador foraneo que apunta a la tabla tipo_contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_general.lugar_ejecucion IS 'identificador foraneo de la informacion del lugar de ejecucion del contrato';
-- ddl-end --
COMMENT ON CONSTRAINT ck_numero_contrato_contrato_general ON argo.contrato_general IS 'check que restringe el formato de numero_contrato  a numero y a tres letras (DVE|DPL|RES|FPL) como prefijo para
los casos de contratación especial';
-- ddl-end --
ALTER TABLE argo.contrato_general OWNER TO postgres;
-- ddl-end --

-- object: argo.contrato | type: TABLE --
CREATE TABLE argo.contrato(
	id_contrato_normal integer NOT NULL DEFAULT nextval('argo.contrato_id_contrato_seq'::regclass),
	vigencia integer NOT NULL,
	numero_contrato character varying NOT NULL,
	estado_registro boolean NOT NULL DEFAULT true,
	fecha_registro date NOT NULL,
	numero_convenio integer,
	vigencia_convenio integer,
	perfil character varying,
	CONSTRAINT pk_contrato PRIMARY KEY (id_contrato_normal)

);
-- ddl-end --
-- object: fki_contrato_general_contrato | type: INDEX --
CREATE INDEX fki_contrato_general_contrato ON argo.contrato
	USING btree
	(
	  numero_contrato,
	  vigencia
	)	WITH (FILLFACTOR = 90);
-- ddl-end --


COMMENT ON COLUMN argo.contrato.perfil IS 'perfil del contratista';
-- ddl-end --
ALTER TABLE argo.contrato OWNER TO postgres;
-- ddl-end --

-- object: argo.contrato_estado | type: TABLE --
CREATE TABLE argo.contrato_estado(
	numero_contrato character varying,
	vigencia integer,
	fecha_registro timestamp,
	id integer NOT NULL DEFAULT nextval('argo.consecutivo_unico_cambio_estado_seq'::regclass),
	estado integer,
	usuario character varying,
	CONSTRAINT pk_contrato_estado PRIMARY KEY (id)

);
-- ddl-end --
-- object: fki_contrato_general_contrato_estado | type: INDEX --
CREATE INDEX fki_contrato_general_contrato_estado ON argo.contrato_estado
	USING btree
	(
	  numero_contrato,
	  vigencia
	)	WITH (FILLFACTOR = 90);
-- ddl-end --

-- object: fki_estado_contrato_contrato_estado | type: INDEX --
CREATE INDEX fki_estado_contrato_contrato_estado ON argo.contrato_estado
	USING btree
	(
	  estado
	)	WITH (FILLFACTOR = 90);
-- ddl-end --


COMMENT ON TABLE argo.contrato_estado IS 'tabla encargada de registrar la relacion entre los contratos y los estados.';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_estado.numero_contrato IS 'numero del contrato (consecutivo Unico)';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_estado.vigencia IS 'año de vigencia del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_estado.fecha_registro IS 'fecha de registro del cambio de estado';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_estado.id IS 'identificador unico del registro de cambio de algun estado ';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_estado.estado IS 'identificador foraneo que referencia a la tabla estado_contrato e indica el estado del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.contrato_estado.usuario IS 'usuario que realiza el cambio de estado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_contrato_estado ON argo.contrato_estado IS 'llave primaria de la tabla contrato_estado es una secuencia.';
-- ddl-end --
ALTER TABLE argo.contrato_estado OWNER TO postgres;
-- ddl-end --

-- object: argo.lugar_ejecucion | type: TABLE --
CREATE TABLE argo.lugar_ejecucion(
	id integer NOT NULL DEFAULT nextval('argo.lugar_ejecucion_id_seq'::regclass),
	direccion character varying NOT NULL,
	sede character varying,
	dependencia character varying,
	ciudad numeric(10,0) NOT NULL,
	CONSTRAINT pl_lugar_ejecucion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE argo.lugar_ejecucion IS 'tabla que almacena el lugar de ejecucion del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.lugar_ejecucion.id IS 'identificador unico del regitro';
-- ddl-end --
COMMENT ON COLUMN argo.lugar_ejecucion.direccion IS 'direccion del lugar de ejecucion del contrato';
-- ddl-end --
COMMENT ON COLUMN argo.lugar_ejecucion.sede IS 'sede de ejecucion del contrato, puede ser nula en caso de que el contrato no se ejecute en ninguna sede de la universidad';
-- ddl-end --
COMMENT ON COLUMN argo.lugar_ejecucion.dependencia IS 'dependencia, puede ser nula, similar a la sede';
-- ddl-end --
COMMENT ON COLUMN argo.lugar_ejecucion.ciudad IS 'ciudad donde se ejecuta el contrato';
-- ddl-end --
ALTER TABLE argo.lugar_ejecucion OWNER TO postgres;
-- ddl-end --

-- object: argo.grupo_tipo_contrato | type: TABLE --
CREATE TABLE argo.grupo_tipo_contrato(
	id integer NOT NULL DEFAULT nextval('argo.grupo_tipo_contrato_id_seq'::regclass),
	grupo_contrato character varying NOT NULL,
	descripcion character varying NOT NULL,
	codigo_contraloria character varying,
	CONSTRAINT pk_grupo_tipo_contrato PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE argo.grupo_tipo_contrato IS 'clasificacion de los tipos especificos del contrato ';
-- ddl-end --
COMMENT ON COLUMN argo.grupo_tipo_contrato.id IS 'identificador unico del registro';
-- ddl-end --
COMMENT ON COLUMN argo.grupo_tipo_contrato.grupo_contrato IS 'nombre del grupo de tipo de contrato';
-- ddl-end --
COMMENT ON COLUMN argo.grupo_tipo_contrato.descripcion IS 'descripcion del grupo de tipo de contrato';
-- ddl-end --
COMMENT ON COLUMN argo.grupo_tipo_contrato.codigo_contraloria IS 'codigo de la contraloria que aplica para el tipo de compromiso en el reporte de sivicof pero que en el entorno de argo se usa como el grupo de tipologias';
-- ddl-end --
ALTER TABLE argo.grupo_tipo_contrato OWNER TO postgres;
-- ddl-end --

-- object: core.arl | type: TABLE --
CREATE TABLE core.arl(
	nit numeric(20,0) NOT NULL,
	id_ubicacion numeric(10,0),
	nombre character varying(25) NOT NULL,
	direccion character varying(30) NOT NULL,
	telefono numeric(10,0) NOT NULL,
	extencion_telefono numeric(4,0),
	fax numeric(10,0),
	extencion_fax numeric(4,0),
	nombre_representante_legal character varying(40),
	email character varying(30) NOT NULL,
	estado character varying(8) NOT NULL,
	lugar numeric(5,0),
	CONSTRAINT ckc_estado_arl CHECK (((estado)::text = ANY (ARRAY[('Activo'::character varying)::text, ('Inactivo'::character varying)::text]))),
	CONSTRAINT pk_arl PRIMARY KEY (nit)

);
-- ddl-end --
-- object: arl_pk | type: INDEX --
CREATE UNIQUE INDEX arl_pk ON core.arl
	USING btree
	(
	  nit
	)	WITH (FILLFACTOR = 90);
-- ddl-end --


ALTER TABLE core.arl OWNER TO postgres;
-- ddl-end --

-- object: core.caja_compensacion | type: TABLE --
CREATE TABLE core.caja_compensacion(
	nit numeric(20,0) NOT NULL,
	id_ubicacion numeric(10,0),
	nombre character varying(25) NOT NULL,
	direccion character varying(30) NOT NULL,
	telefono numeric(10,0) NOT NULL,
	extencion_telefono numeric(4,0),
	fax numeric(10,0),
	extencion_fax numeric(5,0),
	nombre_representante_legal character varying(40),
	email character varying(30) NOT NULL,
	estado character varying(8) NOT NULL,
	CONSTRAINT ckc_estado_caja_com CHECK (((estado)::text = ANY (ARRAY[('Activo'::character varying)::text, ('Inactivo'::character varying)::text]))),
	CONSTRAINT pk_caja_compensacion PRIMARY KEY (nit)

);
-- ddl-end --
-- object: caja_compensacion_pk | type: INDEX --
CREATE UNIQUE INDEX caja_compensacion_pk ON core.caja_compensacion
	USING btree
	(
	  nit
	)	WITH (FILLFACTOR = 90);
-- ddl-end --


ALTER TABLE core.caja_compensacion OWNER TO postgres;
-- ddl-end --

-- object: core.eps | type: TABLE --
CREATE TABLE core.eps(
	nit numeric(20,0) NOT NULL,
	id_ubicacion numeric(10,0),
	nombre character varying(25) NOT NULL,
	direccion character varying(30) NOT NULL,
	telefono numeric(10,0) NOT NULL,
	ext_telefono numeric(4,0),
	fax numeric(10,0),
	ext_fax numeric(4,0),
	nom_representante character varying(40),
	email character varying(30) NOT NULL,
	estado character varying(8) NOT NULL,
	CONSTRAINT ckc_estado_eps CHECK (((estado)::text = ANY (ARRAY[('Activo'::character varying)::text, ('Inactivo'::character varying)::text]))),
	CONSTRAINT pk_eps PRIMARY KEY (nit)

);
-- ddl-end --
-- object: eps_pk | type: INDEX --
CREATE UNIQUE INDEX eps_pk ON core.eps
	USING btree
	(
	  nit
	)	WITH (FILLFACTOR = 90);
-- ddl-end --


ALTER TABLE core.eps OWNER TO postgres;
-- ddl-end --

-- object: core.fondo_pension | type: TABLE --
CREATE TABLE core.fondo_pension(
	nit numeric(15,0) NOT NULL,
	id_ubicacion numeric(10,0),
	nombre character varying(25) NOT NULL,
	direccion character varying(30) NOT NULL,
	telefono numeric(10,0) NOT NULL,
	ext_telefono numeric(4,0),
	fax numeric(10,0),
	ext_fax numeric(4,0),
	nom_representante character varying(40),
	email character varying(30) NOT NULL,
	estado character varying(8) NOT NULL,
	CONSTRAINT pk_fondo_pension PRIMARY KEY (nit)

);
-- ddl-end --
-- object: fondo_pensiones_pk | type: INDEX --
CREATE UNIQUE INDEX fondo_pensiones_pk ON core.fondo_pension
	USING btree
	(
	  nit
	)	WITH (FILLFACTOR = 90);
-- ddl-end --


ALTER TABLE core.fondo_pension OWNER TO postgres;
-- ddl-end --

-- object: personal.resolucion | type: TABLE --
CREATE TABLE personal.resolucion(
	id integer NOT NULL DEFAULT nextval('resolucion_id_seq'::regclass),
	num_resolucion text NOT NULL,
	fecha_emision date NOT NULL,
	vigencia text NOT NULL,
	id_facultad integer NOT NULL,
	CONSTRAINT pk_id_resolucion PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE personal.resolucion OWNER TO postgres;
-- ddl-end --

-- object: public.clasificacion | type: TABLE --
CREATE TABLE public.clasificacion(
	id serial,
	fecha_registro timestamp with time zone NOT NULL,
	id_categoria integer NOT NULL,
	id_dedicacion integer,
	id_docente integer NOT NULL,
	CONSTRAINT pk_clasificacion PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.categoria | type: TABLE --
CREATE TABLE public.categoria(
	id integer,
	nombre character varying(20) NOT NULL,
	descripcion character varying(500),
	CONSTRAINT pk_categoria PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.dedicacion | type: TABLE --
CREATE TABLE public.dedicacion(
	id integer,
	nombre character varying(20) NOT NULL,
	descripcion character varying(500),
	CONSTRAINT pk_dedicacion PRIMARY KEY (id)

);
-- ddl-end --
-- object: categoria_fk | type: CONSTRAINT --
ALTER TABLE public.clasificacion ADD CONSTRAINT categoria_fk FOREIGN KEY (id_categoria)
REFERENCES public.categoria (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: dedicacion_fk | type: CONSTRAINT --
ALTER TABLE public.clasificacion ADD CONSTRAINT dedicacion_fk FOREIGN KEY (id_dedicacion)
REFERENCES public.dedicacion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: public.docente | type: TABLE --
CREATE TABLE public.docente(
	id serial,
	codigo character varying(12) NOT NULL,
	num_documento_persona_informacion_persona_natural numeric(15,0),
	CONSTRAINT pk_docente PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.contrato_especial | type: TABLE --
CREATE TABLE public.contrato_especial(
	id serial,
	id_resolucion integer DEFAULT nextval('resolucion_id_seq'::regclass),
	id_docente integer NOT NULL,
	id_contrato_normal_contrato integer DEFAULT nextval('argo.contrato_id_contrato_seq'::regclass),
	tipo integer NOT NULL,
	CONSTRAINT pk_contrato_especial PRIMARY KEY (id)

);
-- ddl-end --
-- object: docente_fk | type: CONSTRAINT --
ALTER TABLE public.clasificacion ADD CONSTRAINT docente_fk FOREIGN KEY (id_docente)
REFERENCES public.docente (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: informacion_persona_natural_fk | type: CONSTRAINT --
ALTER TABLE public.docente ADD CONSTRAINT informacion_persona_natural_fk FOREIGN KEY (num_documento_persona_informacion_persona_natural)
REFERENCES agora.informacion_persona_natural (num_documento_persona) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: docente_uq | type: CONSTRAINT --
ALTER TABLE public.docente ADD CONSTRAINT docente_uq UNIQUE (num_documento_persona_informacion_persona_natural);
-- ddl-end --


-- object: public.proyecto_curricular | type: TABLE --
CREATE TABLE public.proyecto_curricular(
	id serial NOT NULL,
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	id_facultad integer NOT NULL,
	id_tipo_carrera integer,
	CONSTRAINT pk_proyecto_curricular PRIMARY KEY (id)

);
-- ddl-end --
-- object: public.facultad | type: TABLE --
CREATE TABLE public.facultad(
	id serial NOT NULL,
	nombre character varying(100) NOT NULL,
	direccion character varying(100),
	CONSTRAINT pk_facultad PRIMARY KEY (id)

);
-- ddl-end --
-- object: facultad_fk | type: CONSTRAINT --
ALTER TABLE personal.resolucion ADD CONSTRAINT facultad_fk FOREIGN KEY (id_facultad)
REFERENCES public.facultad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: resolucion_fk | type: CONSTRAINT --
ALTER TABLE public.contrato_especial ADD CONSTRAINT resolucion_fk FOREIGN KEY (id_resolucion)
REFERENCES personal.resolucion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: docente_fk | type: CONSTRAINT --
ALTER TABLE public.contrato_especial ADD CONSTRAINT docente_fk FOREIGN KEY (id_docente)
REFERENCES public.docente (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: contrato_fk | type: CONSTRAINT --
ALTER TABLE public.contrato_especial ADD CONSTRAINT contrato_fk FOREIGN KEY (id_contrato_normal_contrato)
REFERENCES argo.contrato (id_contrato_normal) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: contrato_especial_uq | type: CONSTRAINT --
ALTER TABLE public.contrato_especial ADD CONSTRAINT contrato_especial_uq UNIQUE (id_contrato_normal_contrato);
-- ddl-end --


-- object: public.carga_academica | type: TABLE --
CREATE TABLE public.carga_academica(
	id serial,
	horas_semanales integer NOT NULL,
	vigencia integer NOT NULL,
	periodo integer NOT NULL,
	id_docente integer,
	id_proyecto_curricular integer,
	CONSTRAINT pk_carga_academica PRIMARY KEY (id)

);
-- ddl-end --
-- object: facultad_fk | type: CONSTRAINT --
ALTER TABLE public.proyecto_curricular ADD CONSTRAINT facultad_fk FOREIGN KEY (id_facultad)
REFERENCES public.facultad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: docente_fk | type: CONSTRAINT --
ALTER TABLE public.carga_academica ADD CONSTRAINT docente_fk FOREIGN KEY (id_docente)
REFERENCES public.docente (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: proyecto_curricular_fk | type: CONSTRAINT --
ALTER TABLE public.carga_academica ADD CONSTRAINT proyecto_curricular_fk FOREIGN KEY (id_proyecto_curricular)
REFERENCES public.proyecto_curricular (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: public.tipo_carrera | type: TABLE --
CREATE TABLE public.tipo_carrera(
	id serial,
	nombre character varying(20) NOT NULL,
	"descripción" character varying(500),
	CONSTRAINT pk_tipo_carrera PRIMARY KEY (id)

);
-- ddl-end --
-- object: tipo_carrera_fk | type: CONSTRAINT --
ALTER TABLE public.proyecto_curricular ADD CONSTRAINT tipo_carrera_fk FOREIGN KEY (id_tipo_carrera)
REFERENCES public.tipo_carrera (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: "fki_idDepartamento" | type: CONSTRAINT --
ALTER TABLE core.ciudad ADD CONSTRAINT "fki_idDepartamento" FOREIGN KEY (id_departamento)
REFERENCES core.departamento (id_departamento) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: "fki_idPais" | type: CONSTRAINT --
ALTER TABLE core.departamento ADD CONSTRAINT "fki_idPais" FOREIGN KEY (id_pais)
REFERENCES core.pais (id_pais) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: id_perfil | type: CONSTRAINT --
ALTER TABLE agora.informacion_persona_natural ADD CONSTRAINT id_perfil FOREIGN KEY (perfil)
REFERENCES agora.parametro_estandar (id_parametro) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: id_tipo_documento | type: CONSTRAINT --
ALTER TABLE agora.informacion_persona_natural ADD CONSTRAINT id_tipo_documento FOREIGN KEY (tipo_documento)
REFERENCES agora.parametro_estandar (id_parametro) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: "idNucleo" | type: CONSTRAINT --
ALTER TABLE agora.informacion_persona_natural ADD CONSTRAINT "idNucleo" FOREIGN KEY (id_nucleo_basico)
REFERENCES core.snies_nucleo_basico (id_nucleo) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: id_pais_nacimiento | type: CONSTRAINT --
ALTER TABLE agora.informacion_persona_natural ADD CONSTRAINT id_pais_nacimiento FOREIGN KEY (id_pais_nacimiento)
REFERENCES core.pais (id_pais) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: fki_id_banco_afc | type: CONSTRAINT --
ALTER TABLE agora.informacion_persona_natural ADD CONSTRAINT fki_id_banco_afc FOREIGN KEY (id_entidad_bancaria_afc)
REFERENCES core.banco (id_codigo) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: fki_id_arl | type: CONSTRAINT --
ALTER TABLE agora.informacion_persona_natural ADD CONSTRAINT fki_id_arl FOREIGN KEY (id_arl)
REFERENCES core.parametro_entidad (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: fki_id_caja_compensacion | type: CONSTRAINT --
ALTER TABLE agora.informacion_persona_natural ADD CONSTRAINT fki_id_caja_compensacion FOREIGN KEY (id_caja_compensacion)
REFERENCES core.parametro_entidad (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: fki_id_eps | type: CONSTRAINT --
ALTER TABLE agora.informacion_persona_natural ADD CONSTRAINT fki_id_eps FOREIGN KEY (id_eps)
REFERENCES core.parametro_entidad (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: fki_id_fondo_pension | type: CONSTRAINT --
ALTER TABLE agora.informacion_persona_natural ADD CONSTRAINT fki_id_fondo_pension FOREIGN KEY (id_fondo_pension)
REFERENCES core.parametro_entidad (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: fki_nit_arl | type: CONSTRAINT --
ALTER TABLE agora.informacion_persona_natural ADD CONSTRAINT fki_nit_arl FOREIGN KEY (id_nit_arl)
REFERENCES core.arl (nit) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: fki_nit_eps | type: CONSTRAINT --
ALTER TABLE agora.informacion_persona_natural ADD CONSTRAINT fki_nit_eps FOREIGN KEY (id_nit_eps)
REFERENCES core.eps (nit) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: fki_nit_fondo_pension | type: CONSTRAINT --
ALTER TABLE agora.informacion_persona_natural ADD CONSTRAINT fki_nit_fondo_pension FOREIGN KEY (id_nit_fondo_pension)
REFERENCES core.fondo_pension (nit) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: fki_nit_caja_compensacion | type: CONSTRAINT --
ALTER TABLE agora.informacion_persona_natural ADD CONSTRAINT fki_nit_caja_compensacion FOREIGN KEY (id_nit_caja_compensacion)
REFERENCES core.caja_compensacion (nit) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE CASCADE NOT DEFERRABLE;
-- ddl-end --


-- object: fk_grupo_tipo_contrato_tipo_contrato | type: CONSTRAINT --
ALTER TABLE argo.tipo_contrato ADD CONSTRAINT fk_grupo_tipo_contrato_tipo_contrato FOREIGN KEY (id_grupo_tipo_contrato)
REFERENCES argo.grupo_tipo_contrato (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: fk_lugar_ejecucion_contrato_general | type: CONSTRAINT --
ALTER TABLE argo.contrato_general ADD CONSTRAINT fk_lugar_ejecucion_contrato_general FOREIGN KEY (lugar_ejecucion)
REFERENCES argo.lugar_ejecucion (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: fk_parametro_forma_pago | type: CONSTRAINT --
ALTER TABLE argo.contrato_general ADD CONSTRAINT fk_parametro_forma_pago FOREIGN KEY (forma_pago)
REFERENCES argo.parametros (id_parametro) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: fk_parametro_unidad_tiempo_ejecucion_parametros | type: CONSTRAINT --
ALTER TABLE argo.contrato_general ADD CONSTRAINT fk_parametro_unidad_tiempo_ejecucion_parametros FOREIGN KEY (unidad_ejecucion)
REFERENCES argo.parametros (id_parametro) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: "fk_tipo_contrato_contrato_generaL" | type: CONSTRAINT --
ALTER TABLE argo.contrato_general ADD CONSTRAINT "fk_tipo_contrato_contrato_generaL" FOREIGN KEY (tipo_contrato)
REFERENCES argo.tipo_contrato (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: fk_unidad_ejecutora_contrato_general | type: CONSTRAINT --
ALTER TABLE argo.contrato_general ADD CONSTRAINT fk_unidad_ejecutora_contrato_general FOREIGN KEY (unidad_ejecutora)
REFERENCES kronos.unidad_ejecutora (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: fk_supervisor_contrato_contrato_general | type: CONSTRAINT --
ALTER TABLE argo.contrato_general ADD CONSTRAINT fk_supervisor_contrato_contrato_general FOREIGN KEY (supervisor)
REFERENCES argo.supervisor_contrato (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: fk_contrato_general_contrato | type: CONSTRAINT --
ALTER TABLE argo.contrato ADD CONSTRAINT fk_contrato_general_contrato FOREIGN KEY (numero_contrato,vigencia)
REFERENCES argo.contrato_general (numero_contrato,vigencia) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --


-- object: fk_contrato_general_contrato_estado | type: CONSTRAINT --
ALTER TABLE argo.contrato_estado ADD CONSTRAINT fk_contrato_general_contrato_estado FOREIGN KEY (numero_contrato,vigencia)
REFERENCES argo.contrato_general (numero_contrato,vigencia) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION NOT DEFERRABLE;
-- ddl-end --



