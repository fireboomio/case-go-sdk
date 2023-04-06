package generated

type (
    ChatGPT__Chat__CreateOneChatMessageInput struct {
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
    }        
    InternalChatGPT__Chat__CreateOneChatMessageInput struct {
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
    }        
    InjectedChatGPT__Chat__CreateOneChatMessageInput struct {
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
    }        
    ChatGPT__Chat__CreateOneChatMessageResponse struct {
        Data ChatGPT__Chat__CreateOneChatMessageResponseData `json:"data"`
    }
    ChatGPT__Chat__CreateOneChatMessageResponseData struct {
        Sqlite_createOneUntitled1 ChatGPT__Chat__CreateOneChatMessageResponseData_Sqlite_createOneUntitled1 `json:"sqlite_createOneUntitled1,omitempty"`    
    }            
    ChatGPT__Chat__CreateOneChatMessageResponseData_Sqlite_createOneUntitled1 struct {
        CreatedAt string `json:"createdAt,omitempty"`
        DeletedAt string `json:"deletedAt,omitempty"`
    }
)

type (
    ChatGPT__Chat__CreateOneHistoryInput struct {
        Title string `json:"title"`
        Uuid  string `json:"uuid"`
    }        
    InternalChatGPT__Chat__CreateOneHistoryInput struct {
        Title     string `json:"title"`
        UpdatedAt string `json:"updatedAt"`
        Uuid      string `json:"uuid"`
    }        
    InjectedChatGPT__Chat__CreateOneHistoryInput struct {
        Title     string `json:"title"`
        UpdatedAt string `json:"updatedAt"`
        Uuid      string `json:"uuid"`
    }        
    ChatGPT__Chat__CreateOneHistoryResponse struct {
        Data ChatGPT__Chat__CreateOneHistoryResponseData `json:"data"`
    }
    ChatGPT__Chat__CreateOneHistoryResponseData struct {
        Data ChatGPT__Chat__CreateOneHistoryResponseData_Data `json:"data,omitempty"`    
    }            
    ChatGPT__Chat__CreateOneHistoryResponseData_Data struct {
        Id        int `json:"id,omitempty"`
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
        Uuid      string `json:"uuid,omitempty"`
        CreatedAt string `json:"createdAt,omitempty"`
    }
)

type (
    ChatGPT__Chat__DeleteOneChatMessageInput struct {
        Id int `json:"id"`
    }        
    InternalChatGPT__Chat__DeleteOneChatMessageInput struct {
        Id int `json:"id"`
    }        
    InjectedChatGPT__Chat__DeleteOneChatMessageInput struct {
        Id int `json:"id"`
    }        
    ChatGPT__Chat__DeleteOneChatMessageResponse struct {
        Data ChatGPT__Chat__DeleteOneChatMessageResponseData `json:"data"`
    }
    ChatGPT__Chat__DeleteOneChatMessageResponseData struct {
        Data ChatGPT__Chat__DeleteOneChatMessageResponseData_Data `json:"data,omitempty"`    
    }            
    ChatGPT__Chat__DeleteOneChatMessageResponseData_Data struct {
        Id int `json:"id,omitempty"`
    }
)

type (
    ChatGPT__Chat__DeleteOneHistoryInput struct {
        Id int `json:"id"`
    }        
    InternalChatGPT__Chat__DeleteOneHistoryInput struct {
        Id int `json:"id"`
    }        
    InjectedChatGPT__Chat__DeleteOneHistoryInput struct {
        Id int `json:"id"`
    }        
    ChatGPT__Chat__DeleteOneHistoryResponse struct {
        Data ChatGPT__Chat__DeleteOneHistoryResponseData `json:"data"`
    }
    ChatGPT__Chat__DeleteOneHistoryResponseData struct {
        Data ChatGPT__Chat__DeleteOneHistoryResponseData_Data `json:"data,omitempty"`    
    }            
    ChatGPT__Chat__DeleteOneHistoryResponseData_Data struct {
        Id int `json:"id,omitempty"`
    }
)

type (
    ChatGPT__Chat__GetHistoryListInput struct {
        Skip    int `json:"skip,omitempty"`
        Take    int `json:"take,omitempty"`
        OrderBy []*Grace_HistoryOrderByWithRelationInput `json:"orderBy,omitempty"`                
        Query   *Grace_HistoryWhereInput `json:"query,omitempty"`
    }        
    InternalChatGPT__Chat__GetHistoryListInput struct {
        Take    int `json:"take,omitempty"`
        OrderBy []*Grace_HistoryOrderByWithRelationInput `json:"orderBy,omitempty"`                
        Query   *Grace_HistoryWhereInput `json:"query,omitempty"`
        Skip    int `json:"skip,omitempty"`
    }        
    Grace_NestedStringFilter struct {
        Lt         string `json:"lt,omitempty"`
        Lte        string `json:"lte,omitempty"`
        Not        *Grace_NestedStringFilter `json:"not,omitempty"`
        StartsWith string `json:"startsWith,omitempty"`
        EndsWith   string `json:"endsWith,omitempty"`
        Equals     string `json:"equals,omitempty"`
        Gt         string `json:"gt,omitempty"`
        NotIn      []string `json:"notIn,omitempty"`                
        Contains   string `json:"contains,omitempty"`
        Gte        string `json:"gte,omitempty"`
        In         []string `json:"in,omitempty"`                
    }    
    Grace_StringFilter struct {
        Lte        string `json:"lte,omitempty"`
        Not        *Grace_NestedStringFilter `json:"not,omitempty"`
        StartsWith string `json:"startsWith,omitempty"`
        Contains   string `json:"contains,omitempty"`
        EndsWith   string `json:"endsWith,omitempty"`
        Gte        string `json:"gte,omitempty"`
        In         []string `json:"in,omitempty"`                
        Lt         string `json:"lt,omitempty"`
        Equals     string `json:"equals,omitempty"`
        Gt         string `json:"gt,omitempty"`
        NotIn      []string `json:"notIn,omitempty"`                
    }    
    Grace_NestedBigIntFilter struct {
        Lt     string `json:"lt,omitempty"`
        Lte    string `json:"lte,omitempty"`
        Not    *Grace_NestedBigIntFilter `json:"not,omitempty"`
        NotIn  []string `json:"notIn,omitempty"`                
        Equals string `json:"equals,omitempty"`
        Gt     string `json:"gt,omitempty"`
        Gte    string `json:"gte,omitempty"`
        In     []string `json:"in,omitempty"`                
    }    
    Grace_NestedIntFilter struct {
        NotIn  []int `json:"notIn,omitempty"`                
        Equals int `json:"equals,omitempty"`
        Gt     int `json:"gt,omitempty"`
        Gte    int `json:"gte,omitempty"`
        In     []int `json:"in,omitempty"`                
        Lt     int `json:"lt,omitempty"`
        Lte    int `json:"lte,omitempty"`
        Not    *Grace_NestedIntFilter `json:"not,omitempty"`
    }    
    Grace_DateTimeNullableFilter struct {
        In     []string `json:"in,omitempty"`                
        Lt     string `json:"lt,omitempty"`
        Lte    string `json:"lte,omitempty"`
        Not    *Grace_NestedDateTimeNullableFilter `json:"not,omitempty"`
        NotIn  []string `json:"notIn,omitempty"`                
        Equals string `json:"equals,omitempty"`
        Gt     string `json:"gt,omitempty"`
        Gte    string `json:"gte,omitempty"`
    }    
    Grace_HistoryOrderByWithRelationInput struct {
        UpdatedAt string `json:"updatedAt,omitempty"`
        Uuid      string `json:"uuid,omitempty"`
        CreatedAt string `json:"createdAt,omitempty"`
        Id        string `json:"id,omitempty"`
        Title     string `json:"title,omitempty"`
    }    
    Grace_HistoryWhereInput struct {
        UpdatedAt *Grace_DateTimeNullableFilter `json:"updatedAt,omitempty"`
        Uuid      *Grace_BigIntFilter `json:"uuid,omitempty"`
        AND       *Grace_HistoryWhereInput `json:"AND,omitempty"`
        NOT       *Grace_HistoryWhereInput `json:"NOT,omitempty"`
        OR        []*Grace_HistoryWhereInput `json:"OR,omitempty"`                
        CreatedAt *Grace_DateTimeFilter `json:"createdAt,omitempty"`
        Id        *Grace_IntFilter `json:"id,omitempty"`
        Title     *Grace_StringFilter `json:"title,omitempty"`
    }    
    Grace_IntFilter struct {
        Gt     int `json:"gt,omitempty"`
        Gte    int `json:"gte,omitempty"`
        In     []int `json:"in,omitempty"`                
        Lt     int `json:"lt,omitempty"`
        Lte    int `json:"lte,omitempty"`
        Not    *Grace_NestedIntFilter `json:"not,omitempty"`
        NotIn  []int `json:"notIn,omitempty"`                
        Equals int `json:"equals,omitempty"`
    }    
    Grace_NestedDateTimeFilter struct {
        Gte    string `json:"gte,omitempty"`
        In     []string `json:"in,omitempty"`                
        Lt     string `json:"lt,omitempty"`
        Lte    string `json:"lte,omitempty"`
        Not    *Grace_NestedDateTimeFilter `json:"not,omitempty"`
        NotIn  []string `json:"notIn,omitempty"`                
        Equals string `json:"equals,omitempty"`
        Gt     string `json:"gt,omitempty"`
    }    
    Grace_NestedDateTimeNullableFilter struct {
        Lt     string `json:"lt,omitempty"`
        Lte    string `json:"lte,omitempty"`
        Not    *Grace_NestedDateTimeNullableFilter `json:"not,omitempty"`
        NotIn  []string `json:"notIn,omitempty"`                
        Equals string `json:"equals,omitempty"`
        Gt     string `json:"gt,omitempty"`
        Gte    string `json:"gte,omitempty"`
        In     []string `json:"in,omitempty"`                
    }    
    Grace_BigIntFilter struct {
        Lte    string `json:"lte,omitempty"`
        Not    *Grace_NestedBigIntFilter `json:"not,omitempty"`
        NotIn  []string `json:"notIn,omitempty"`                
        Equals string `json:"equals,omitempty"`
        Gt     string `json:"gt,omitempty"`
        Gte    string `json:"gte,omitempty"`
        In     []string `json:"in,omitempty"`                
        Lt     string `json:"lt,omitempty"`
    }    
    Grace_DateTimeFilter struct {
        Gt     string `json:"gt,omitempty"`
        Gte    string `json:"gte,omitempty"`
        In     []string `json:"in,omitempty"`                
        Lt     string `json:"lt,omitempty"`
        Lte    string `json:"lte,omitempty"`
        Not    *Grace_NestedDateTimeFilter `json:"not,omitempty"`
        NotIn  []string `json:"notIn,omitempty"`                
        Equals string `json:"equals,omitempty"`
    }    
    InjectedChatGPT__Chat__GetHistoryListInput struct {
        Query   *Grace_HistoryWhereInput `json:"query,omitempty"`
        Skip    int `json:"skip,omitempty"`
        Take    int `json:"take,omitempty"`
        OrderBy []*Grace_HistoryOrderByWithRelationInput `json:"orderBy,omitempty"`                
    }        
    ChatGPT__Chat__GetHistoryListResponse struct {
        Data ChatGPT__Chat__GetHistoryListResponseData `json:"data"`
    }
    ChatGPT__Chat__GetHistoryListResponseData struct {
        Data  []ChatGPT__Chat__GetHistoryListResponseData_Data `json:"data"`    
        Total int `json:"total"`
    }            
    ChatGPT__Chat__GetHistoryListResponseData_Data struct {
        CreatedAt string `json:"createdAt,omitempty"`
        Id        int `json:"id,omitempty"`
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
        Uuid      string `json:"uuid,omitempty"`
    }
)

type (
    ChatGPT__Chat__UpdateOneHistoryInput struct {
        Id    int `json:"id"`
        Title string `json:"title,omitempty"`
        Uuid  string `json:"uuid,omitempty"`
    }        
    InternalChatGPT__Chat__UpdateOneHistoryInput struct {
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt"`
        Uuid      string `json:"uuid,omitempty"`
        Id        int `json:"id"`
    }        
    InjectedChatGPT__Chat__UpdateOneHistoryInput struct {
        Uuid      string `json:"uuid,omitempty"`
        Id        int `json:"id"`
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt"`
    }        
    ChatGPT__Chat__UpdateOneHistoryResponse struct {
        Data ChatGPT__Chat__UpdateOneHistoryResponseData `json:"data"`
    }
    ChatGPT__Chat__UpdateOneHistoryResponseData struct {
        Data ChatGPT__Chat__UpdateOneHistoryResponseData_Data `json:"data,omitempty"`    
    }            
    ChatGPT__Chat__UpdateOneHistoryResponseData_Data struct {
        Uuid      string `json:"uuid,omitempty"`
        CreatedAt string `json:"createdAt,omitempty"`
        Id        int `json:"id,omitempty"`
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
    }
)

type (
    ChatGPT__Propmt__CreateOnePromptInput struct {
        Prompt string `json:"prompt"`
        Title  string `json:"title"`
    }        
    InternalChatGPT__Propmt__CreateOnePromptInput struct {
        Prompt    string `json:"prompt"`
        Title     string `json:"title"`
        UpdatedAt string `json:"updatedAt"`
    }        
    InjectedChatGPT__Propmt__CreateOnePromptInput struct {
        Prompt    string `json:"prompt"`
        Title     string `json:"title"`
        UpdatedAt string `json:"updatedAt"`
    }        
    ChatGPT__Propmt__CreateOnePromptResponse struct {
        Data ChatGPT__Propmt__CreateOnePromptResponseData `json:"data"`
    }
    ChatGPT__Propmt__CreateOnePromptResponseData struct {
        Data ChatGPT__Propmt__CreateOnePromptResponseData_Data `json:"data,omitempty"`    
    }            
    ChatGPT__Propmt__CreateOnePromptResponseData_Data struct {
        Id        int `json:"id,omitempty"`
        Prompt    string `json:"prompt,omitempty"`
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
        CreatedAt string `json:"createdAt,omitempty"`
    }
)

type (
    ChatGPT__Propmt__DeleteManyPromptInput struct {
        Ids []int `json:"ids"`                
    }        
    InternalChatGPT__Propmt__DeleteManyPromptInput struct {
        Ids []int `json:"ids"`                
    }        
    InjectedChatGPT__Propmt__DeleteManyPromptInput struct {
        Ids []int `json:"ids"`                
    }        
    ChatGPT__Propmt__DeleteManyPromptResponse struct {
        Data ChatGPT__Propmt__DeleteManyPromptResponseData `json:"data"`
    }
    ChatGPT__Propmt__DeleteManyPromptResponseData struct {
        Data ChatGPT__Propmt__DeleteManyPromptResponseData_Data `json:"data,omitempty"`    
    }            
    ChatGPT__Propmt__DeleteManyPromptResponseData_Data struct {
        Count int `json:"count,omitempty"`
    }
)

type (
    ChatGPT__Propmt__DeleteOnePromptInput struct {
        Id int `json:"id"`
    }        
    InternalChatGPT__Propmt__DeleteOnePromptInput struct {
        Id int `json:"id"`
    }        
    InjectedChatGPT__Propmt__DeleteOnePromptInput struct {
        Id int `json:"id"`
    }        
    ChatGPT__Propmt__DeleteOnePromptResponse struct {
        Data ChatGPT__Propmt__DeleteOnePromptResponseData `json:"data"`
    }
    ChatGPT__Propmt__DeleteOnePromptResponseData struct {
        Data ChatGPT__Propmt__DeleteOnePromptResponseData_Data `json:"data,omitempty"`    
    }            
    ChatGPT__Propmt__DeleteOnePromptResponseData_Data struct {
        Id int `json:"id,omitempty"`
    }
)

type (
    ChatGPT__Propmt__GetPromptListInput struct {
        OrderBy []*Grace_PromptOrderByWithRelationInput `json:"orderBy,omitempty"`                
        Query   *Grace_PromptWhereInput `json:"query,omitempty"`
        Skip    int `json:"skip,omitempty"`
        Take    int `json:"take,omitempty"`
    }        
    InternalChatGPT__Propmt__GetPromptListInput struct {
        OrderBy []*Grace_PromptOrderByWithRelationInput `json:"orderBy,omitempty"`                
        Query   *Grace_PromptWhereInput `json:"query,omitempty"`
        Skip    int `json:"skip,omitempty"`
        Take    int `json:"take,omitempty"`
    }        

    Grace_PromptWhereInput struct {
        AND       *Grace_PromptWhereInput `json:"AND,omitempty"`
        NOT       *Grace_PromptWhereInput `json:"NOT,omitempty"`
        OR        []*Grace_PromptWhereInput `json:"OR,omitempty"`                
        CreatedAt *Grace_DateTimeFilter `json:"createdAt,omitempty"`
        Id        *Grace_IntFilter `json:"id,omitempty"`
        Prompt    *Grace_StringFilter `json:"prompt,omitempty"`
        Title     *Grace_StringFilter `json:"title,omitempty"`
        UpdatedAt *Grace_DateTimeNullableFilter `json:"updatedAt,omitempty"`
    }    





    Grace_PromptOrderByWithRelationInput struct {
        Prompt    string `json:"prompt,omitempty"`
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
        CreatedAt string `json:"createdAt,omitempty"`
        Id        string `json:"id,omitempty"`
    }    


    InjectedChatGPT__Propmt__GetPromptListInput struct {
        Query   *Grace_PromptWhereInput `json:"query,omitempty"`
        Skip    int `json:"skip,omitempty"`
        Take    int `json:"take,omitempty"`
        OrderBy []*Grace_PromptOrderByWithRelationInput `json:"orderBy,omitempty"`                
    }        
    ChatGPT__Propmt__GetPromptListResponse struct {
        Data ChatGPT__Propmt__GetPromptListResponseData `json:"data"`
    }
    ChatGPT__Propmt__GetPromptListResponseData struct {
        Data  []ChatGPT__Propmt__GetPromptListResponseData_Data `json:"data"`    
        Total int `json:"total"`
    }            
    ChatGPT__Propmt__GetPromptListResponseData_Data struct {
        CreatedAt string `json:"createdAt,omitempty"`
        Id        int `json:"id,omitempty"`
        Prompt    string `json:"prompt,omitempty"`
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
    }
)

type (
    ChatGPT__Propmt__UpdateOnePromptInput struct {
        Title  string `json:"title,omitempty"`
        Id     int `json:"id"`
        Prompt string `json:"prompt,omitempty"`
    }        
    InternalChatGPT__Propmt__UpdateOnePromptInput struct {
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt"`
        Id        int `json:"id"`
        Prompt    string `json:"prompt,omitempty"`
    }        
    InjectedChatGPT__Propmt__UpdateOnePromptInput struct {
        Prompt    string `json:"prompt,omitempty"`
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt"`
        Id        int `json:"id"`
    }        
    ChatGPT__Propmt__UpdateOnePromptResponse struct {
        Data ChatGPT__Propmt__UpdateOnePromptResponseData `json:"data"`
    }
    ChatGPT__Propmt__UpdateOnePromptResponseData struct {
        Data ChatGPT__Propmt__UpdateOnePromptResponseData_Data `json:"data,omitempty"`    
    }            
    ChatGPT__Propmt__UpdateOnePromptResponseData_Data struct {
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
        CreatedAt string `json:"createdAt,omitempty"`
        Id        int `json:"id,omitempty"`
        Prompt    string `json:"prompt,omitempty"`
    }
)

type (
    QQ__WW__EEInput struct {
        Prompt string `json:"prompt,omitempty"`
        Title  string `json:"title,omitempty"`
    }        
    InternalQQ__WW__EEInput struct {
        Title  string `json:"title,omitempty"`
        Prompt string `json:"prompt,omitempty"`
    }        
    InjectedQQ__WW__EEInput struct {
        Prompt string `json:"prompt,omitempty"`
        Title  string `json:"title,omitempty"`
    }        
    QQ__WW__EEResponse struct {
        Data QQ__WW__EEResponseData `json:"data"`
    }
    QQ__WW__EEResponseData struct {
        Grace_createOnePrompt QQ__WW__EEResponseData_Grace_createOnePrompt `json:"grace_createOnePrompt,omitempty"`    
    }            
    QQ__WW__EEResponseData_Grace_createOnePrompt struct {
        Id int `json:"id,omitempty"`
    }
)

type (
    Sqlite__TestInput struct {
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt"`
    }        
    InternalSqlite__TestInput struct {
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt"`
    }        
    InjectedSqlite__TestInput struct {
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt"`
    }        
    Sqlite__TestResponse struct {
        Data Sqlite__TestResponseData `json:"data"`
    }
    Sqlite__TestResponseData struct {
        Sqlite_createOneUntitled1 Sqlite__TestResponseData_Sqlite_createOneUntitled1 `json:"sqlite_createOneUntitled1,omitempty"`    
    }            
    Sqlite__TestResponseData_Sqlite_createOneUntitled1 struct {
        CreatedAt string `json:"createdAt,omitempty"`
        DeletedAt string `json:"deletedAt,omitempty"`
        Id        int `json:"id,omitempty"`
        Title     string `json:"title,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
    }
)

type (
    System__BindRoleApisInput struct {
        AllRoles []string `json:"allRoles"`                
        Apis     []int `json:"apis"`                
        RoleCode string `json:"roleCode"`
    }        
    InternalSystem__BindRoleApisInput struct {
        AllRoles []string `json:"allRoles"`                
        Apis     []int `json:"apis"`                
        RoleCode string `json:"roleCode"`
    }        
    InjectedSystem__BindRoleApisInput struct {
        AllRoles []string `json:"allRoles"`                
        Apis     []int `json:"apis"`                
        RoleCode string `json:"roleCode"`
    }        
    System__BindRoleApisResponse struct {
        Data System__BindRoleApisResponseData `json:"data"`
    }
    System__BindRoleApisResponseData struct {
        Data System__BindRoleApisResponseData_Data `json:"data,omitempty"`    
    }            
    System__BindRoleApisResponseData_Data struct {
        Count int `json:"count,omitempty"`
    }
)

type (
    System__GetRoleBindApisInput struct {
        Code string `json:"code"`
    }        
    InternalSystem__GetRoleBindApisInput struct {
        Code string `json:"code"`
    }        
    InjectedSystem__GetRoleBindApisInput struct {
        Code string `json:"code"`
    }        
    System__GetRoleBindApisResponse struct {
        Data System__GetRoleBindApisResponseData `json:"data"`
    }
    System__GetRoleBindApisResponseData struct {
        Data []System__GetRoleBindApisResponseData_Data `json:"data,omitempty"`    
    }            
    System__GetRoleBindApisResponseData_Data struct {
        Id            int `json:"id,omitempty"`
        Method        string `json:"method,omitempty"`
        OperationType string `json:"operationType,omitempty"`
        Remark        string `json:"remark,omitempty"`
        Roles         string `json:"roles,omitempty"`
        CreateTime    string `json:"createTime,omitempty"`
        DeleteTime    string `json:"deleteTime,omitempty"`
        Enabled       bool `json:"enabled,omitempty"`
        RestUrl       string `json:"restUrl,omitempty"`
        Title         string `json:"title,omitempty"`
        Illegal       bool `json:"illegal,omitempty"`
        IsPublic      bool `json:"isPublic,omitempty"`
        Content       string `json:"content,omitempty"`
        LiveQuery     bool `json:"liveQuery,omitempty"`
        RoleType      string `json:"roleType,omitempty"`
        UpdateTime    string `json:"updateTime,omitempty"`
    }
)

