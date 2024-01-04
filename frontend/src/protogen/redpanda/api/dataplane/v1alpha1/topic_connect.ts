// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/dataplane/v1alpha1/topic.proto (package redpanda.api.dataplane.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateTopicRequest, CreateTopicResponse, DeleteTopicRequest, DeleteTopicResponse, GetTopicConfigurationRequest, GetTopicConfigurationResponse, ListTopicsRequest, ListTopicsResponse, SetTopicConfigurationRequest, SetTopicConfigurationResponse, UpdateTopicConfigurationRequest, UpdateTopicConfigurationResponse } from "./topic_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service redpanda.api.dataplane.v1alpha1.TopicService
 */
export const TopicService = {
  typeName: "redpanda.api.dataplane.v1alpha1.TopicService",
  methods: {
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.TopicService.CreateTopic
     */
    createTopic: {
      name: "CreateTopic",
      I: CreateTopicRequest,
      O: CreateTopicResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.TopicService.ListTopics
     */
    listTopics: {
      name: "ListTopics",
      I: ListTopicsRequest,
      O: ListTopicsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.TopicService.DeleteTopic
     */
    deleteTopic: {
      name: "DeleteTopic",
      I: DeleteTopicRequest,
      O: DeleteTopicResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.TopicService.GetTopicConfiguration
     */
    getTopicConfiguration: {
      name: "GetTopicConfiguration",
      I: GetTopicConfigurationRequest,
      O: GetTopicConfigurationResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.TopicService.UpdateTopicConfiguration
     */
    updateTopicConfiguration: {
      name: "UpdateTopicConfiguration",
      I: UpdateTopicConfigurationRequest,
      O: UpdateTopicConfigurationResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1alpha1.TopicService.SetTopicConfiguration
     */
    setTopicConfiguration: {
      name: "SetTopicConfiguration",
      I: SetTopicConfigurationRequest,
      O: SetTopicConfigurationResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
