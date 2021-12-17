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
proto.proto = require('./service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.crudServiceClient =
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
proto.proto.crudServicePromiseClient =
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
 *   !proto.proto.CreateRequest,
 *   !proto.proto.AuthResponse>}
 */
const methodDescriptor_crudService_Create = new grpc.web.MethodDescriptor(
  '/proto.crudService/Create',
  grpc.web.MethodType.UNARY,
  proto.proto.CreateRequest,
  proto.proto.AuthResponse,
  /**
   * @param {!proto.proto.CreateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.AuthResponse.deserializeBinary
);


/**
 * @param {!proto.proto.CreateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.AuthResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.AuthResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.crudServiceClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.crudService/Create',
      request,
      metadata || {},
      methodDescriptor_crudService_Create,
      callback);
};


/**
 * @param {!proto.proto.CreateRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.AuthResponse>}
 *     Promise that resolves to the response
 */
proto.proto.crudServicePromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.crudService/Create',
      request,
      metadata || {},
      methodDescriptor_crudService_Create);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.RetrieveRequest,
 *   !proto.proto.AuthResponse>}
 */
const methodDescriptor_crudService_Retrieve = new grpc.web.MethodDescriptor(
  '/proto.crudService/Retrieve',
  grpc.web.MethodType.UNARY,
  proto.proto.RetrieveRequest,
  proto.proto.AuthResponse,
  /**
   * @param {!proto.proto.RetrieveRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.AuthResponse.deserializeBinary
);


/**
 * @param {!proto.proto.RetrieveRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.AuthResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.AuthResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.crudServiceClient.prototype.retrieve =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.crudService/Retrieve',
      request,
      metadata || {},
      methodDescriptor_crudService_Retrieve,
      callback);
};


/**
 * @param {!proto.proto.RetrieveRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.AuthResponse>}
 *     Promise that resolves to the response
 */
proto.proto.crudServicePromiseClient.prototype.retrieve =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.crudService/Retrieve',
      request,
      metadata || {},
      methodDescriptor_crudService_Retrieve);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.UpdateRequest,
 *   !proto.proto.AuthResponse>}
 */
const methodDescriptor_crudService_Update = new grpc.web.MethodDescriptor(
  '/proto.crudService/Update',
  grpc.web.MethodType.UNARY,
  proto.proto.UpdateRequest,
  proto.proto.AuthResponse,
  /**
   * @param {!proto.proto.UpdateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.AuthResponse.deserializeBinary
);


/**
 * @param {!proto.proto.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.AuthResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.AuthResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.crudServiceClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.crudService/Update',
      request,
      metadata || {},
      methodDescriptor_crudService_Update,
      callback);
};


/**
 * @param {!proto.proto.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.AuthResponse>}
 *     Promise that resolves to the response
 */
proto.proto.crudServicePromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.crudService/Update',
      request,
      metadata || {},
      methodDescriptor_crudService_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.DeleteRequest,
 *   !proto.proto.AuthResponse>}
 */
const methodDescriptor_crudService_Delete = new grpc.web.MethodDescriptor(
  '/proto.crudService/Delete',
  grpc.web.MethodType.UNARY,
  proto.proto.DeleteRequest,
  proto.proto.AuthResponse,
  /**
   * @param {!proto.proto.DeleteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.AuthResponse.deserializeBinary
);


/**
 * @param {!proto.proto.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.AuthResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.AuthResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.crudServiceClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.crudService/Delete',
      request,
      metadata || {},
      methodDescriptor_crudService_Delete,
      callback);
};


/**
 * @param {!proto.proto.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.AuthResponse>}
 *     Promise that resolves to the response
 */
proto.proto.crudServicePromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.crudService/Delete',
      request,
      metadata || {},
      methodDescriptor_crudService_Delete);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.EidUsedRequest,
 *   !proto.proto.UsedResponse>}
 */
const methodDescriptor_crudService_EidUsed = new grpc.web.MethodDescriptor(
  '/proto.crudService/EidUsed',
  grpc.web.MethodType.UNARY,
  proto.proto.EidUsedRequest,
  proto.proto.UsedResponse,
  /**
   * @param {!proto.proto.EidUsedRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.UsedResponse.deserializeBinary
);


/**
 * @param {!proto.proto.EidUsedRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.UsedResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.UsedResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.crudServiceClient.prototype.eidUsed =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.crudService/EidUsed',
      request,
      metadata || {},
      methodDescriptor_crudService_EidUsed,
      callback);
};


/**
 * @param {!proto.proto.EidUsedRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.UsedResponse>}
 *     Promise that resolves to the response
 */
proto.proto.crudServicePromiseClient.prototype.eidUsed =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.crudService/EidUsed',
      request,
      metadata || {},
      methodDescriptor_crudService_EidUsed);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.AuthUserRequest,
 *   !proto.proto.AuthUserResponse>}
 */
const methodDescriptor_crudService_AuthUser = new grpc.web.MethodDescriptor(
  '/proto.crudService/AuthUser',
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
proto.proto.crudServiceClient.prototype.authUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.crudService/AuthUser',
      request,
      metadata || {},
      methodDescriptor_crudService_AuthUser,
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
proto.proto.crudServicePromiseClient.prototype.authUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.crudService/AuthUser',
      request,
      metadata || {},
      methodDescriptor_crudService_AuthUser);
};


module.exports = proto.proto;

