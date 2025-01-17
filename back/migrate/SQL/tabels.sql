-- Active: 1721397412112@@127.0.0.1@5432@softwhere
 

CREATE TABLE IF NOT EXISTS public.GOODS ( 
    GOODS_ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    REF     VARCHAR(1024) , 
    Photo   TEXT
);

CREATE TABLE IF NOT EXISTS public.BOUGHT ( 
    BOUGHT_ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    GOODS_ID UUID NOT NULL, 
    quantity INTEGER, 
    FOREIGN KEY (GOODS_ID) REFERENCES public.GOODS(GOODS_ID)
);

CREATE TABLE IF NOT EXISTS public.STATES ( 
    STATE_ID  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    state_name NAME   
); 

CREATE TABLE IF NOT EXISTS public.Transaction ( 
    TRANS_ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    USER_ID UUID , 
    STATE_ID  UUID ,
    Photo PATH , 
    FOREIGN KEY (STATE_ID) REFERENCES public.STATES(STATE_ID) , 
    FOREIGN KEY (USER_ID) REFERENCES public.USERS(ID)  
);


CREATE TABLE IF NOT EXISTS public.Transaction_BOUGHT ( 
    BOUGHT_ID UUID,
    TRANS_ID UUID, 
    Photo PATH, 
    FOREIGN KEY (BOUGHT_ID) REFERENCES public.BOUGHT(BOUGHT_ID), 
    FOREIGN KEY (TRANS_ID) REFERENCES public.Transaction(TRANS_ID) 
);


