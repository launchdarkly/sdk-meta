/* eslint-disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

/**
 * Defines LaunchDarkly SDK Log Codes
 */
export interface LogCodes {
  systems: {
    [k: string]: System;
  };
  classes: {
    [k: string]: Class;
  };
  conditions: {
    [k: string]: Condition;
  };
  [k: string]: unknown;
}
export interface System {
  specifier: number;
  description: string;
}
export interface Class {
  specifier: number;
  description: string;
  additionalProperties?: never;
  [k: string]: unknown;
}
export interface Condition {
  specifier: number;
  class: number;
  system: number;
  description: string;
  message: Message;
  deprecated?: boolean;
  deprecatedReason?: string;
  superseded?: string;
  supersededReason?: string;
}
export interface Message {
  parameterized: string;
  parameters?: {
    [k: string]: string;
  };
}
