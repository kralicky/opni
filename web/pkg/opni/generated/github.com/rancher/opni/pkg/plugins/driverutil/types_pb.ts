// @generated by protoc-gen-es v1.3.0 with parameter "target=ts,import_extension=none,ts_nocheck=false"
// @generated from file github.com/rancher/opni/pkg/plugins/driverutil/types.proto (package driverutil, syntax proto3)
/* eslint-disable */

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Revision } from "../../apis/core/v1/core_pb";

/**
 * @generated from enum driverutil.Target
 */
export enum Target {
  /**
   * @generated from enum value: ActiveConfiguration = 0;
   */
  ActiveConfiguration = 0,

  /**
   * @generated from enum value: DefaultConfiguration = 1;
   */
  DefaultConfiguration = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(Target)
proto3.util.setEnumType(Target, "driverutil.Target", [
  { no: 0, name: "ActiveConfiguration" },
  { no: 1, name: "DefaultConfiguration" },
]);

/**
 * @generated from enum driverutil.Action
 */
export enum Action {
  /**
   * @generated from enum value: NoAction = 0;
   */
  NoAction = 0,

  /**
   * @generated from enum value: Set = 1;
   */
  Set = 1,

  /**
   * @generated from enum value: Reset = 2;
   */
  Reset = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(Action)
proto3.util.setEnumType(Action, "driverutil.Action", [
  { no: 0, name: "NoAction" },
  { no: 1, name: "Set" },
  { no: 2, name: "Reset" },
]);

/**
 * @generated from enum driverutil.InstallState
 */
export enum InstallState {
  /**
   * @generated from enum value: NotInstalled = 0;
   */
  NotInstalled = 0,

  /**
   * @generated from enum value: Installed = 1;
   */
  Installed = 1,

  /**
   * @generated from enum value: Uninstalling = 2;
   */
  Uninstalling = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(InstallState)
proto3.util.setEnumType(InstallState, "driverutil.InstallState", [
  { no: 0, name: "NotInstalled" },
  { no: 1, name: "Installed" },
  { no: 2, name: "Uninstalling" },
]);

/**
 * @generated from enum driverutil.ConfigurationState
 */
export enum ConfigurationState {
  /**
   * @generated from enum value: NotConfigured = 0;
   */
  NotConfigured = 0,

  /**
   * @generated from enum value: Configured = 1;
   */
  Configured = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(ConfigurationState)
proto3.util.setEnumType(ConfigurationState, "driverutil.ConfigurationState", [
  { no: 0, name: "NotConfigured" },
  { no: 1, name: "Configured" },
]);

/**
 * @generated from enum driverutil.ApplicationState
 */
export enum ApplicationState {
  /**
   * @generated from enum value: NotRunning = 0;
   */
  NotRunning = 0,

  /**
   * @generated from enum value: Pending = 1;
   */
  Pending = 1,

  /**
   * @generated from enum value: Running = 2;
   */
  Running = 2,

  /**
   * @generated from enum value: Failed = 3;
   */
  Failed = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(ApplicationState)
proto3.util.setEnumType(ApplicationState, "driverutil.ApplicationState", [
  { no: 0, name: "NotRunning" },
  { no: 1, name: "Pending" },
  { no: 2, name: "Running" },
  { no: 3, name: "Failed" },
]);

/**
 * Get request options. See also: [pkg/storage.GetOptions]
 *
 * @generated from message driverutil.GetRequest
 */
export class GetRequest extends Message<GetRequest> {
  /**
   * If set, will return the config at the specified revision instead of
   * the current config.
   * This revision value can be obtained from the revision field of a
   * typed GetConfiguration/GetDefaultConfiguration response, or from one of
   * the history entries in a typed ConfigurationHistory response.
   *
   * @generated from field: core.Revision revision = 1;
   */
  revision?: Revision;

  constructor(data?: PartialMessage<GetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "driverutil.GetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "revision", kind: "message", T: Revision },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRequest {
    return new GetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRequest {
    return new GetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRequest {
    return new GetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetRequest | PlainMessage<GetRequest> | undefined, b: GetRequest | PlainMessage<GetRequest> | undefined): boolean {
    return proto3.util.equals(GetRequest, a, b);
  }
}

/**
 * History request options. See also: [pkg/storage.HistoryOptions]
 *
 * @generated from message driverutil.ConfigurationHistoryRequest
 */
export class ConfigurationHistoryRequest extends Message<ConfigurationHistoryRequest> {
  /**
   * The configuration type to return history for.
   *
   * @generated from field: driverutil.Target target = 1;
   */
  target = Target.ActiveConfiguration;

  /**
   * The latest modification revision to include in the returned history.
   *
   * @generated from field: core.Revision revision = 2;
   */
  revision?: Revision;

  /**
   * If set, will include the values of the configuration in the response.
   * Otherwise, only the revision field of each entry will be populated.
   *
   * @generated from field: bool includeValues = 3;
   */
  includeValues = false;

  constructor(data?: PartialMessage<ConfigurationHistoryRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "driverutil.ConfigurationHistoryRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "target", kind: "enum", T: proto3.getEnumType(Target) },
    { no: 2, name: "revision", kind: "message", T: Revision },
    { no: 3, name: "includeValues", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConfigurationHistoryRequest {
    return new ConfigurationHistoryRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConfigurationHistoryRequest {
    return new ConfigurationHistoryRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConfigurationHistoryRequest {
    return new ConfigurationHistoryRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ConfigurationHistoryRequest | PlainMessage<ConfigurationHistoryRequest> | undefined, b: ConfigurationHistoryRequest | PlainMessage<ConfigurationHistoryRequest> | undefined): boolean {
    return proto3.util.equals(ConfigurationHistoryRequest, a, b);
  }
}

/**
 * @generated from message driverutil.ValidationError
 */
export class ValidationError extends Message<ValidationError> {
  /**
   * @generated from field: driverutil.ValidationError.Severity severity = 1;
   */
  severity = ValidationError_Severity.Unknown;

  /**
   * @generated from field: string message = 2;
   */
  message = "";

  /**
   * @generated from field: string source = 3;
   */
  source = "";

  constructor(data?: PartialMessage<ValidationError>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "driverutil.ValidationError";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "severity", kind: "enum", T: proto3.getEnumType(ValidationError_Severity) },
    { no: 2, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "source", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ValidationError {
    return new ValidationError().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ValidationError {
    return new ValidationError().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ValidationError {
    return new ValidationError().fromJsonString(jsonString, options);
  }

  static equals(a: ValidationError | PlainMessage<ValidationError> | undefined, b: ValidationError | PlainMessage<ValidationError> | undefined): boolean {
    return proto3.util.equals(ValidationError, a, b);
  }
}

/**
 * @generated from enum driverutil.ValidationError.Severity
 */
export enum ValidationError_Severity {
  /**
   * @generated from enum value: Unknown = 0;
   */
  Unknown = 0,

  /**
   * @generated from enum value: Warning = 1;
   */
  Warning = 1,

  /**
   * @generated from enum value: Error = 2;
   */
  Error = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(ValidationError_Severity)
proto3.util.setEnumType(ValidationError_Severity, "driverutil.ValidationError.Severity", [
  { no: 0, name: "Unknown" },
  { no: 1, name: "Warning" },
  { no: 2, name: "Error" },
]);

/**
 * @generated from message driverutil.InstallStatus
 */
export class InstallStatus extends Message<InstallStatus> {
  /**
   * @generated from field: driverutil.ConfigurationState configState = 1;
   */
  configState = ConfigurationState.NotConfigured;

  /**
   * @generated from field: driverutil.InstallState installState = 2;
   */
  installState = InstallState.NotInstalled;

  /**
   * @generated from field: driverutil.ApplicationState appState = 3;
   */
  appState = ApplicationState.NotRunning;

  /**
   * @generated from field: repeated string warnings = 4;
   */
  warnings: string[] = [];

  /**
   * @generated from field: string version = 5;
   */
  version = "";

  /**
   * @generated from field: map<string, string> metadata = 6;
   */
  metadata: { [key: string]: string } = {};

  constructor(data?: PartialMessage<InstallStatus>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "driverutil.InstallStatus";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "configState", kind: "enum", T: proto3.getEnumType(ConfigurationState) },
    { no: 2, name: "installState", kind: "enum", T: proto3.getEnumType(InstallState) },
    { no: 3, name: "appState", kind: "enum", T: proto3.getEnumType(ApplicationState) },
    { no: 4, name: "warnings", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "metadata", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InstallStatus {
    return new InstallStatus().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InstallStatus {
    return new InstallStatus().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InstallStatus {
    return new InstallStatus().fromJsonString(jsonString, options);
  }

  static equals(a: InstallStatus | PlainMessage<InstallStatus> | undefined, b: InstallStatus | PlainMessage<InstallStatus> | undefined): boolean {
    return proto3.util.equals(InstallStatus, a, b);
  }
}

/**
 * @generated from message driverutil.PresetMetadata
 */
export class PresetMetadata extends Message<PresetMetadata> {
  /**
   * A display name for the preset.
   *
   * @generated from field: string displayName = 1;
   */
  displayName = "";

  /**
   * Longer description of the preset.
   *
   * @generated from field: string description = 2;
   */
  description = "";

  /**
   * Optional list of messages that should be displayed to the user when the
   * preset is used (determined at the client's discretion).
   * For example, this may contain additional information about requirements
   * or next steps that the user should be aware of.
   *
   * @generated from field: repeated string notes = 3;
   */
  notes: string[] = [];

  constructor(data?: PartialMessage<PresetMetadata>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "driverutil.PresetMetadata";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "displayName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "notes", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PresetMetadata {
    return new PresetMetadata().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PresetMetadata {
    return new PresetMetadata().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PresetMetadata {
    return new PresetMetadata().fromJsonString(jsonString, options);
  }

  static equals(a: PresetMetadata | PlainMessage<PresetMetadata> | undefined, b: PresetMetadata | PlainMessage<PresetMetadata> | undefined): boolean {
    return proto3.util.equals(PresetMetadata, a, b);
  }
}
