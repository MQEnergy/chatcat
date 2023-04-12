import {Configuration, OpenAIApi} from "openai";

/**
 * 问答非流式返回
 * @param apiKey
 * @param modelName
 * @param prompt
 * @returns {Promise<void>}
 * @constructor
 */
export async function CompletionsNoStream(apiKey, modelName, prompt) {
  const configuration = new Configuration({
    apiKey: apiKey,
  });
  delete configuration.baseOptions.headers['User-Agent'];
  const openai = new OpenAIApi(configuration);
  const res = await openai.createCompletion({
    model: modelName,
    prompt: prompt,
    max_tokens: 1000,
    temperature: 0
  });
  return res
}

export async function CompletionsStream(wsUrl, apiKey, modelName, prompt) {
  
}