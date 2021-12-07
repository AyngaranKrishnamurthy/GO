/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.proto = require('./services_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.EmpServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.EmpServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.CreateEmpRequest,
 *   !proto.proto.CreateEmpResponse>}
 */
const methodDescriptor_EmpService_CreateEmp = new grpc.web.MethodDescriptor(
  '/proto.EmpService/CreateEmp',
  grpc.web.MethodType.UNARY,
  proto.proto.CreateEmpRequest,
  proto.proto.CreateEmpResponse,
  /**
   * @param {!proto.proto.CreateEmpRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.CreateEmpResponse.deserializeBinary
);


/**
 * @param {!proto.proto.CreateEmpRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.CreateEmpResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.CreateEmpResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.EmpServiceClient.prototype.createEmp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.EmpService/CreateEmp',
      request,
      metadata || {},
      methodDescriptor_EmpService_CreateEmp,
      callback);
};


/**
 * @param {!proto.proto.CreateEmpRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.CreateEmpResponse>}
 *     Promise that resolves to the response
 */
proto.proto.EmpServicePromiseClient.prototype.createEmp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.EmpService/CreateEmp',
      request,
      metadata || {},
      methodDescriptor_EmpService_CreateEmp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.ReadEmpRequest,
 *   !proto.proto.ReadEmpResponse>}
 */
const methodDescriptor_EmpService_ReadEmp = new grpc.web.MethodDescriptor(
  '/proto.EmpService/ReadEmp',
  grpc.web.MethodType.UNARY,
  proto.proto.ReadEmpRequest,
  proto.proto.ReadEmpResponse,
  /**
   * @param {!proto.proto.ReadEmpRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.ReadEmpResponse.deserializeBinary
);


/**
 * @param {!proto.proto.ReadEmpRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.ReadEmpResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.ReadEmpResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.EmpServiceClient.prototype.readEmp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.EmpService/ReadEmp',
      request,
      metadata || {},
      methodDescriptor_EmpService_ReadEmp,
      callback);
};


/**
 * @param {!proto.proto.ReadEmpRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.ReadEmpResponse>}
 *     Promise that resolves to the response
 */
proto.proto.EmpServicePromiseClient.prototype.readEmp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.EmpService/ReadEmp',
      request,
      metadata || {},
      methodDescriptor_EmpService_ReadEmp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.UpdateEmpRequest,
 *   !proto.proto.UpdateEmpResponse>}
 */
const methodDescriptor_EmpService_UpdateEmp = new grpc.web.MethodDescriptor(
  '/proto.EmpService/UpdateEmp',
  grpc.web.MethodType.UNARY,
  proto.proto.UpdateEmpRequest,
  proto.proto.UpdateEmpResponse,
  /**
   * @param {!proto.proto.UpdateEmpRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.UpdateEmpResponse.deserializeBinary
);


/**
 * @param {!proto.proto.UpdateEmpRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.UpdateEmpResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.UpdateEmpResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.EmpServiceClient.prototype.updateEmp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.EmpService/UpdateEmp',
      request,
      metadata || {},
      methodDescriptor_EmpService_UpdateEmp,
      callback);
};


/**
 * @param {!proto.proto.UpdateEmpRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.UpdateEmpResponse>}
 *     Promise that resolves to the response
 */
proto.proto.EmpServicePromiseClient.prototype.updateEmp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.EmpService/UpdateEmp',
      request,
      metadata || {},
      methodDescriptor_EmpService_UpdateEmp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.DeleteEmpRequest,
 *   !proto.proto.DeleteEmpResponse>}
 */
const methodDescriptor_EmpService_DeleteEmp = new grpc.web.MethodDescriptor(
  '/proto.EmpService/DeleteEmp',
  grpc.web.MethodType.UNARY,
  proto.proto.DeleteEmpRequest,
  proto.proto.DeleteEmpResponse,
  /**
   * @param {!proto.proto.DeleteEmpRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.DeleteEmpResponse.deserializeBinary
);


/**
 * @param {!proto.proto.DeleteEmpRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.DeleteEmpResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.DeleteEmpResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.EmpServiceClient.prototype.deleteEmp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.EmpService/DeleteEmp',
      request,
      metadata || {},
      methodDescriptor_EmpService_DeleteEmp,
      callback);
};


/**
 * @param {!proto.proto.DeleteEmpRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.DeleteEmpResponse>}
 *     Promise that resolves to the response
 */
proto.proto.EmpServicePromiseClient.prototype.deleteEmp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.EmpService/DeleteEmp',
      request,
      metadata || {},
      methodDescriptor_EmpService_DeleteEmp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.AuthUserRequest,
 *   !proto.proto.AuthUserResponse>}
 */
const methodDescriptor_EmpService_AuthUser = new grpc.web.MethodDescriptor(
  '/proto.EmpService/AuthUser',
  grpc.web.MethodType.UNARY,
  proto.proto.AuthUserRequest,
  proto.proto.AuthUserResponse,
  /**
   * @param {!proto.proto.AuthUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.AuthUserResponse.deserializeBinary
);


/**
 * @param {!proto.proto.AuthUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.AuthUserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.AuthUserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.EmpServiceClient.prototype.authUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.EmpService/AuthUser',
      request,
      metadata || {},
      methodDescriptor_EmpService_AuthUser,
      callback);
};


/**
 * @param {!proto.proto.AuthUserRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.AuthUserResponse>}
 *     Promise that resolves to the response
 */
proto.proto.EmpServicePromiseClient.prototype.authUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.EmpService/AuthUser',
      request,
      metadata || {},
      methodDescriptor_EmpService_AuthUser);
};


module.exports = proto.proto;

