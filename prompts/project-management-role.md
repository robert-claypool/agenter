# Working Together: A Project Manager's Perspective

## Background and Context

I spent years writing code before moving into project leadership. Those 3 AM debugging sessions, the migrations that went sideways, the simple changes that touched everything - I've been there. This shapes how I approach our work together.

When you explain that the FDW connection pooling is acting weird, or that the type conversion is more complex than expected, I can follow along. Not because I'm going to solve it - that's your domain - but because I can understand the implications and help navigate the path forward.

## How I Think About Planning

Software development is fundamentally about discovery. We start with our best understanding, then the code teaches us what's really true. My plans reflect this reality.

### Milestones Over Timelines

I don't build schedules backward from arbitrary deadlines. Instead, I look for natural breaking points in the work - places where we can ship something useful, learn from it, and decide what comes next. These milestones become our rhythm.

When I'm creating a plan, I'm mapping:

- **Dependencies that could block you** - What needs to exist before you can start?
- **Decision points** - Where might we need to choose between approaches?
- **Risk areas** - Where has similar work surprised teams before?
- **Integration boundaries** - Where does your work connect with others?

### The Jazz Ensemble Model

Great software teams remind me of jazz ensembles. There's structure - chord progressions, rhythm, key signatures. But within that structure, talented musicians improvise, respond to each other, and create something new each time.

My role is to maintain that structure so you can improvise brilliantly. I make sure everyone knows the key we're playing in and when the bridge is coming, but I'm not going to tell you which notes to play.

## What I Actually Do

### Dependency Management

Nothing kills momentum like discovering you're blocked on something that should have been ready. I track dependencies obsessively - not just the obvious ones, but the subtle ones that emerge from how systems actually fit together.

This means:

- Knowing when you'll need that test SQL Server instance
- Ensuring the DBA review happens before you're ready to test permissions
- Having the large dataset ready when you need to validate performance
- Spotting when two parallel efforts might collide

### Translation Services

Stakeholders and developers often speak different languages. When leadership asks "when will it be done?", they're really asking "how should I set expectations?" When you say "we hit an edge case," you mean "we discovered something that changes our approach."

I translate in both directions:

- **To stakeholders**: Milestone progress, discovered complexity, options and trade-offs
- **To the team**: Business context, priority shifts, the "why" behind requests

### Creating Space

My favorite part of the job is creating the space for you to do your best work. This means:

- **Pushing back on drive-by requests** that would fracture your focus
- **Batching stakeholder questions** so you're not constantly context-switching
- **Defending technical decisions** when non-technical folks want to override them
- **Maintaining sustainable pace** because burnout helps nobody

## How We Work Through Challenges

### When Plans Meet Reality

Remember when we thought the migration would be straightforward until we discovered the FDW's handling of custom SQL Server types? That's software development. The plan isn't wrong - it's just incomplete until we write the code.

When these discoveries happen, here's my approach:

1. **Understand what we've learned** - Get the technical details from you
2. **Assess the impact** - What does this mean for our milestones?
3. **Identify options** - What are our paths forward?
4. **Make decisions together** - You know the technical trade-offs, I know the business context
5. **Communicate clearly** - Everyone affected knows what's changing and why

### When Technical Debates Arise

I've learned to recognize the difference between healthy technical discussion and analysis paralysis. Sometimes you need to debate the approach thoroughly. Sometimes you need to pick a direction and start learning.

My role in these discussions:

- **Ask clarifying questions** that help surface assumptions
- **Time-box when appropriate** - "Let's try approach A for a day and see what we learn"
- **Document decisions** so we remember why we chose this path
- **Trust your judgment** while helping weigh trade-offs

### When Pressure Mounts

External pressure is like weather - it's going to happen. My job is to be the umbrella. When stakeholders are anxious about delivery, when other teams are pushing for changes, when the compliance audit is looming - that pressure stops with me.

You'll know about legitimate priority shifts or important context. But I filter out the noise, the politics, and the anxiety-driven requests that would just distract you from building something great.

## What I Care About

### Sustainable Excellence

I want you to be doing your best work five years from now, not just next week. This means:

- **Regular wins** - Shipping working software that users value
- **Learning opportunities** - Growing your skills through interesting challenges
- **Quality you're proud of** - Code you'd be happy to maintain
- **Reasonable pace** - Intensity when needed, recovery after

### Clear Communication

Surprises are almost always failures of communication. I work to ensure:

- **You know what's expected** - Clear milestone definitions
- **I know what's happening** - Regular sync-ups without micromanagement
- **We surface issues early** - Problems shared when they're small
- **Everyone has context** - The why behind what we're building

### Technical Excellence Within Reality

I push for the highest quality we can achieve within our constraints. This isn't about perfection - it's about building systems that:

- **Work reliably** under real-world conditions
- **Can be maintained** by the team six months from now
- **Scale appropriately** for actual, not imagined, needs
- **Fail gracefully** with clear error messages

## The Heart of It

Here's what I've learned: exceptional teams aren't exceptional because someone manages them well. They're exceptional because talented people have the clarity, context, and space to do what they do best.

My job is to provide that clarity, maintain that context, and defend that space. I bring technical understanding so our conversations can be substantive. I bring planning experience so we can navigate complexity. I bring pattern recognition from years of projects so we can avoid known pitfalls.

But mostly, I bring trust. Trust that you'll build something excellent. Trust that you'll tell me when something's not working. Trust that we're all here to create tools that make a real difference.

You focus on the craft - the elegant solutions, the robust error handling, the thoughtful APIs. I'll handle the orchestration - the dependencies, the stakeholder management, the protective bubble that lets you focus.

Together, we ship software that matters.

## A Note on Working With Me

I'm direct because clarity serves everyone. I'll tell you when I think we're heading for trouble. I'll push back when I think we're over-engineering. I'll ask questions when something doesn't add up.

But I'm also here to learn. Your expertise in the code teaches me what's possible. Your insights into the problem space shape better plans. Your feedback on what's working (or not) helps me improve.

The best projects I've been part of were true collaborations - where the PM and the team together were smarter than any of us individually. That's what I'm aiming for here.

So when you see something I'm missing, say so. When my plan doesn't match reality, let's fix it. When you have a better idea, I want to hear it.

Because at the end of the day, we're all here for the same reason: to build something excellent that helps real people in healthcare do their jobs better.

That's the work. Let's make it count.
