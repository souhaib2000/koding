class NotificationController extends KDObject

  subjectMap = ->

    JStatusUpdate       : "<a href='#'>status</a>"
    JCodeSnip           : "<a href='#'>code snippet</a>"
    JQuestionActivity   : "<a href='#'>question</a>"
    JDiscussion         : "<a href='#'>discussion</a>"
    JLinkActivity       : "<a href='#'>link</a>"
    JPrivateMessage     : "<a href='#'>private message</a>"
    JOpinion            : "<a href='#'>opinion</a>"
    JTutorial           : "<a href='#'>tutorial</a>"
    JComment            : "<a href='#'>comment</a>"
    JReview             : "<a href='#'>review</a>"

  constructor:->

    super

    @getSingleton('mainController').on "AccountChanged", =>
      @off 'NotificationHasArrived'
      @notificationChannel?.close().off()
      @setListeners()

  setListeners:->

    @notificationChannel = KD.remote.subscribe 'notification',
      serviceType : 'notification'
      isExclusive : yes

    @notificationChannel.on 'message', (notification)=>
      @emit "NotificationHasArrived", notification
      if notification.contents
        @emit notification.event, notification.contents
        @prepareNotification notification

  prepareNotification: (notification)->

    # NOTIFICATION SAMPLES

    # 1 - < actor fullname > commented on your < activity type >.
    # 2 - < actor fullname > also commented on the < activity type > that you commented.
    # 3 - < actor fullname > liked your < activity type >.
    # 4 - < actor fullname > sent you a private message.
    # 5 - < actor fullname > replied to your private message.
    # 6 - < actor fullname > also replied to your private message.
    # 7 - Your membership request to < group title > has been approved.
    # 8 - < actor fullname > has requested access to < group title >.
    # 9 - < actor fullname > has invited you to < group title >.
    # 9 - < actor fullname > has joined < group title >.

    options = {}
    {origin, subject, actionType, actorType} = notification.contents

    isMine = if origin?._id and origin._id is KD.whoami()._id then yes else no
    actor = notification.contents[actorType]

    return  unless actor

    KD.remote.cacheable actor.constructorName, actor.id, (err, actorAccount)=>
      KD.remote.api[subject.constructorName].one _id: subject.id, (err, subjectObj)=>

        actorName = "#{actorAccount.profile.firstName} #{actorAccount.profile.lastName}"

        options.title = switch actionType
          when "reply", "opinion"
            if isMine
              switch subject.constructorName
                when "JPrivateMessage"
                  "#{actorName} replied to your #{subjectMap()[subject.constructorName]}."
                else
                  "#{actorName} commented on your #{subjectMap()[subject.constructorName]}."
            else
              switch subject.constructorName
                when "JPrivateMessage"
                  "#{actorName} also replied to your #{subjectMap()[subject.constructorName]}."
                else
                  originatorName   = "#{origin.profile.firstName} #{origin.profile.lastName}"
                  if actorName is originatorName
                    originatorName = "their own"
                    separator      = ""
                  else
                    separator      = "'s"
                  "#{actorName} also commented on #{originatorName}#{separator} #{subjectMap()[subject.constructorName]}."

          when "like"
            "#{actorName} liked your #{subjectMap()[subject.constructorName]}."
          when "newMessage"
            @emit "NewMessageArrived"
            "#{actorName} sent you a #{subjectMap()[subject.constructorName]}."
          when "groupRequestApproved"
            "Your membership request to <a href='#'>#{subjectObj.title}</a> has been approved."
          when "groupAccessRequested"
            "#{actorName} has requested access to <a href='#'>#{subjectObj.title}</a>."
          when "groupInvited"
            "#{actorName} has invited you to <a href='#'>#{subjectObj.title}</a>."
          when "groupJoined"
            "#{actorName} has joined <a href='#'>#{subjectObj.title}</a>."
          else
            if actorType is "follower"
              "#{actorName} started following you."

        if subject
          options.click = ->
            view = @
            if subject.constructorName is "JPrivateMessage"
              appManager.openApplication "Inbox"
            else if subject.constructorName in ["JComment", "JOpinion"]
              KD.remote.api[subject.constructorName].fetchRelated subject.id, (err, post) ->
                KD.getSingleton('router').handleRoute "/Activity/#{post.slug}", state:post
                # appManager.tell "Activity", "createContentDisplay", post
                view.destroy()
            else if subject.constructorName is 'JGroup'
              suffix = ''
              suffix = '/Dashboard' if actionType is 'groupAccessRequested'
              KD.getSingleton('router').handleRoute "/#{subjectObj.slug}#{suffix}"
              view.destroy()
            else
              # appManager.tell "Activity", "createContentDisplay", post
              KD.getSingleton('router').handleRoute "/Activity/#{subjectObj.slug}", state:post
              view.destroy()

        options.type  = actionType or actorType or ''

        @notify options

  notify:(options  = {})->

    options.title or= 'notification arrived'

    notification = new KDNotificationView
      type     : 'tray'
      cssClass : "mini realtime #{options.type}"
      duration : 10000
      showTimer: yes
      title    : "<span></span>#{options.title}"
      content  : options.content  or null

    notification.once 'click', options.click


