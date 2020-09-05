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
proto.proto = require('./addGarment_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.AddServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.AddServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 *   !proto.proto.Item,
 *   !proto.proto.ItemList>}
 */
const methodDescriptor_AddService_AddGarment = new grpc.web.MethodDescriptor(
  '/proto.AddService/AddGarment',
  grpc.web.MethodType.UNARY,
  proto.proto.Item,
  proto.proto.ItemList,
  /**
   * @param {!proto.proto.Item} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.ItemList.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.Item,
 *   !proto.proto.ItemList>}
 */
const methodInfo_AddService_AddGarment = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.ItemList,
  /**
   * @param {!proto.proto.Item} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.ItemList.deserializeBinary
);


/**
 * @param {!proto.proto.Item} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.ItemList)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.ItemList>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.AddServiceClient.prototype.addGarment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.AddService/AddGarment',
      request,
      metadata || {},
      methodDescriptor_AddService_AddGarment,
      callback);
};


/**
 * @param {!proto.proto.Item} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.ItemList>}
 *     Promise that resolves to the response
 */
proto.proto.AddServicePromiseClient.prototype.addGarment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.AddService/AddGarment',
      request,
      metadata || {},
      methodDescriptor_AddService_AddGarment);
};


module.exports = proto.proto;

