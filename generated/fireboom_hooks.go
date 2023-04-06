package generated
import "custom-go/pkg/base"

type (
    ChatGPT__Chat__CreateOneChatMessageBody = *base.OperationBody[ChatGPT__Chat__CreateOneChatMessageInput, ChatGPT__Chat__CreateOneChatMessageResponse]
    ChatGPT__Chat__CreateOneHistoryBody     = *base.OperationBody[ChatGPT__Chat__CreateOneHistoryInput, ChatGPT__Chat__CreateOneHistoryResponse]
    ChatGPT__Chat__DeleteOneChatMessageBody = *base.OperationBody[ChatGPT__Chat__DeleteOneChatMessageInput, ChatGPT__Chat__DeleteOneChatMessageResponse]
    ChatGPT__Chat__DeleteOneHistoryBody     = *base.OperationBody[ChatGPT__Chat__DeleteOneHistoryInput, ChatGPT__Chat__DeleteOneHistoryResponse]
    ChatGPT__Chat__GetHistoryListBody       = *base.OperationBody[ChatGPT__Chat__GetHistoryListInput, ChatGPT__Chat__GetHistoryListResponse]
    ChatGPT__Chat__UpdateOneHistoryBody     = *base.OperationBody[ChatGPT__Chat__UpdateOneHistoryInput, ChatGPT__Chat__UpdateOneHistoryResponse]
    ChatGPT__Propmt__CreateOnePromptBody    = *base.OperationBody[ChatGPT__Propmt__CreateOnePromptInput, ChatGPT__Propmt__CreateOnePromptResponse]
    ChatGPT__Propmt__DeleteManyPromptBody   = *base.OperationBody[ChatGPT__Propmt__DeleteManyPromptInput, ChatGPT__Propmt__DeleteManyPromptResponse]
    ChatGPT__Propmt__DeleteOnePromptBody    = *base.OperationBody[ChatGPT__Propmt__DeleteOnePromptInput, ChatGPT__Propmt__DeleteOnePromptResponse]
    ChatGPT__Propmt__GetPromptListBody      = *base.OperationBody[ChatGPT__Propmt__GetPromptListInput, ChatGPT__Propmt__GetPromptListResponse]
    ChatGPT__Propmt__UpdateOnePromptBody    = *base.OperationBody[ChatGPT__Propmt__UpdateOnePromptInput, ChatGPT__Propmt__UpdateOnePromptResponse]
    QQ__WW__EEBody                          = *base.OperationBody[QQ__WW__EEInput, QQ__WW__EEResponse]
    Sqlite__TestBody                        = *base.OperationBody[Sqlite__TestInput, Sqlite__TestResponse]
    System__BindRoleApisBody                = *base.OperationBody[System__BindRoleApisInput, System__BindRoleApisResponse]
    System__GetRoleBindApisBody             = *base.OperationBody[System__GetRoleBindApisInput, System__GetRoleBindApisResponse]
)

const (
    ChatGPT__Chat__GetHistoryList  base.OperationQueryPath = "ChatGPT/Chat/GetHistoryList"
    ChatGPT__Propmt__GetPromptList base.OperationQueryPath = "ChatGPT/Propmt/GetPromptList"
    System__GetRoleBindApis        base.OperationQueryPath = "System/GetRoleBindApis"
)

const (
    ChatGPT__Chat__CreateOneChatMessage base.OperationMutationPath = "ChatGPT/Chat/CreateOneChatMessage"
    ChatGPT__Chat__CreateOneHistory     base.OperationMutationPath = "ChatGPT/Chat/CreateOneHistory"
    ChatGPT__Chat__DeleteOneChatMessage base.OperationMutationPath = "ChatGPT/Chat/DeleteOneChatMessage"
    ChatGPT__Chat__DeleteOneHistory     base.OperationMutationPath = "ChatGPT/Chat/DeleteOneHistory"
    ChatGPT__Chat__UpdateOneHistory     base.OperationMutationPath = "ChatGPT/Chat/UpdateOneHistory"
    ChatGPT__Propmt__CreateOnePrompt    base.OperationMutationPath = "ChatGPT/Propmt/CreateOnePrompt"
    ChatGPT__Propmt__DeleteManyPrompt   base.OperationMutationPath = "ChatGPT/Propmt/DeleteManyPrompt"
    ChatGPT__Propmt__DeleteOnePrompt    base.OperationMutationPath = "ChatGPT/Propmt/DeleteOnePrompt"
    ChatGPT__Propmt__UpdateOnePrompt    base.OperationMutationPath = "ChatGPT/Propmt/UpdateOnePrompt"
    QQ__WW__EE                          base.OperationMutationPath = "QQ/WW/EE"
    Sqlite__Test                        base.OperationMutationPath = "Sqlite/Test"
    System__BindRoleApis                base.OperationMutationPath = "System/BindRoleApis"
)
