# Comment and Documentation Guidelines

## The Core Principle

Write for expert software engineers who appreciate simple, direct language. They can understand complex concepts but are delighted by clear explanations.

### Key Insights

1. **Simple words aren't dumbing down** - They remove cognitive load so users can focus on what matters.

2. **Don't drop important details** - When changing "Check prerequisites (PostgreSQL, tds_fdw, permissions)" to something simpler, keep the specific tools listed. Don't just say "Check if ready."

3. **Avoid babying the reader** - "Test migration by copying 10 rows from each table" → "Test migration with 10 rows per table". Trust that developers know what "per table" means.

4. **Question every technical term** - Ask: Is this the right word, or am I using it because it sounds professional?
   - "Prerequisites" → "required tools" or "your setup"
   - "Perform migration" → "Copy data"
   - "Execute" → "Run"

5. **Be action-oriented** - Say what the command does, not what category it falls into.

6. **Let narrative drive structure, not headers** - Too many headers is a cheat. If you need lots of headers to organize your thoughts, your prose isn't clear enough. Good narrative flows naturally and only needs headers for visual separation to help maintain reader attention. The document shouldn't be jammed into a rigid outline structure.

## Writing Checklist

Before writing any help text or comment:

1. **Write the first draft naturally** - How would you explain this to a smart colleague?

2. **Check each technical term** - Is there a simpler word that means the same thing?
   - Aggregate → Combine (unless you really mean mathematical aggregation)
   - Utilize → Use
   - Prerequisites → Requirements, needed tools, or just list what's needed

3. **Remove unnecessary words** - But keep the ones that add meaning
   - "Test and save connection" → "Save connection" (testing is implied)
   - "From each table" → "per table"
   - But keep: "(PostgreSQL, tds_fdw)" because it's specific and helpful

4. **Read it out loud** - Does it sound like something a human would say?

## Examples from This Session

**Before:**
```
Check prerequisites (PostgreSQL, tds_fdw, permissions)
Perform full migration  
Test migration by copying 10 rows from each table
```

**After:**
```
Check required tools (PostgreSQL, tds_fdw extension)
Copy all data
Test migration with 10 rows per table
```

**Even better:**
```
Check your setup
Copy all tables from SQL Server to PostgreSQL
Test with sample data (10 rows per table)
```

## The Test

Would an experienced developer:
1. Understand immediately what this does?
2. Appreciate that we didn't waste their time with corporate-speak?
3. Have all the details they need to succeed?

If yes to all three, the comment is good.

## Remember

The best documentation respects the reader's intelligence while removing unnecessary friction. We're not writing for a compliance audit or a process document. We're writing for humans who want to get work done.
